package authservice

import (
	"context"
	"encoding/json"

	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/utils"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func (s *serv) GetAccessToken(ctx context.Context, token string) (string, error) {
	claims, err := utils.VerifyToken(token, s.jwtConfig.RefreshSecretKey())
	if err != nil {
		return "", err
	}

	err = s.checkTokenRefresh(token)
	if err != nil {
		return "", err
	}

	info, err := s.getUserInfoFromStorage(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	if claims.Role != info.Role {
		return "", errors.New("authentication error")
	}

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     claims.Role,
	}, s.jwtConfig.AccessSecretKey(), s.jwtConfig.AccessExpirationTime())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *serv) GetRefreshToken(ctx context.Context, oldToken string) (string, error) {
	claims, err := utils.VerifyToken(oldToken, s.jwtConfig.RefreshSecretKey())
	if err != nil {
		return "", err
	}

	err = s.checkTokenRefresh(oldToken)
	if err != nil {
		return "", err
	}

	info, err := s.getUserInfoFromStorage(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	if claims.Role != info.Role {
		return "", errors.New("authentication error")
	}

	res := s.redis.Set(oldToken, nil, s.jwtConfig.RefreshExpirationTime())
	if res.Err() != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     claims.Role,
	}, s.jwtConfig.RefreshSecretKey(), s.jwtConfig.RefreshExpirationTime())
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *serv) getUserInfoFromStorage(ctx context.Context, username string) (*model.UserInfo, error) {
	var info *model.UserInfo

	res, err := s.redis.Get(username).Result()
	if errors.Is(err, redis.Nil) {
		user, errRep := s.userRepo.GetByUsername(ctx, username)
		if errRep != nil {
			return nil, errRep
		}

		info = &user.Info
	} else if err != nil {
		return nil, err
	}

	if info == nil {
		err = json.Unmarshal([]byte(res), &info)
		if err != nil {
			return nil, err
		}
	}

	return info, nil
}

func (s *serv) checkTokenRefresh(refreshToken string) error {
	_, err := s.redis.Get(refreshToken).Result()
	if errors.Is(err, redis.Nil) {
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("refresh token has expired")
}
