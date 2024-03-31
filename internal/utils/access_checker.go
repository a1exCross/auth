package utils

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/a1exCross/auth/internal/config"
	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/repository"

	"github.com/a1exCross/common/pkg/storage"

	"github.com/go-redis/redis"
)

// AccessChecker - верификатор доступа
type AccessChecker interface {
	RefreshCheck(context.Context, string, ...model.UserRole) (bool, error)
	AccessCheck(context.Context, string, ...model.UserRole) (bool, error)
}

type routeAccessChecker struct {
	jwtConfig config.JWTConfig
	redis     storage.Redis
	userRepo  repository.UserRepository
}

// NewRouteAccessChecker - создает новый экземпляр верификатора
func NewRouteAccessChecker(jwtCfg config.JWTConfig, redis storage.Redis, repo repository.UserRepository) AccessChecker {
	return &routeAccessChecker{
		jwtConfig: jwtCfg,
		redis:     redis,
		userRepo:  repo,
	}
}

// RefreshCheck - проверяет refresh токен на валидность
func (r *routeAccessChecker) RefreshCheck(ctx context.Context, token string, roles ...model.UserRole) (bool, error) {
	if len(roles) == 0 {
		return true, nil
	}

	claims, err := VerifyToken(token, r.jwtConfig.RefreshSecretKey())
	if err != nil {
		return false, err
	}

	var info *model.UserInfo

	res, err := r.redis.Get(claims.Username).Result()
	if errors.Is(err, redis.Nil) {
		user, errRep := r.userRepo.GetByUsername(ctx, claims.Username)
		if errRep != nil {
			return false, errRep
		}

		info = &user.Info
	} else if err != nil {
		return false, err
	}

	if info == nil {
		err = json.Unmarshal([]byte(res), &info)
		if err != nil {
			return false, err
		}
	}

	for _, role := range roles {
		if role == info.Role {
			return true, nil
		}
	}

	return false, nil
}

// AccessCheck - проверяет access токен на валидность
func (r *routeAccessChecker) AccessCheck(ctx context.Context, token string, roles ...model.UserRole) (bool, error) {
	if len(roles) == 0 {
		return true, nil
	}

	claims, err := VerifyToken(token, r.jwtConfig.AccessSecretKey())
	if err != nil {
		return false, err
	}

	var info *model.UserInfo

	cmd := r.redis.Get(claims.Username)

	var res string
	if cmd != nil {
		res, err = cmd.Result()
	}

	if errors.Is(err, redis.Nil) {
		user, errRep := r.userRepo.GetByUsername(ctx, claims.Username)
		if errRep != nil {
			return false, errRep
		}

		info = &user.Info
	} else if err != nil {
		return false, err
	}

	if info == nil {
		err = json.Unmarshal([]byte(res), &info)
		if err != nil {
			return false, err
		}
	}

	if info.Role == model.ADMIN {
		return true, nil
	}

	for _, role := range roles {
		if role == info.Role {
			return true, nil
		}
	}

	return false, nil
}
