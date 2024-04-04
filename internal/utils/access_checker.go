package utils

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/a1exCross/auth/internal/config"
	"github.com/a1exCross/auth/internal/model"
	"github.com/a1exCross/auth/internal/repository"

	"github.com/a1exCross/common/pkg/filter"
	"github.com/a1exCross/common/pkg/storage"

	"github.com/go-redis/redis"
)

// AccessChecker - верификатор доступа
type AccessChecker interface {
	AccessCheck(ctx context.Context, token string, endpoint string) (bool, error)
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

// AccessCheck - проверяет access токен на валидность
func (r *routeAccessChecker) AccessCheck(ctx context.Context, token string, endpoint string) (bool, error) {
	claims, err := VerifyToken(token, r.jwtConfig.AccessSecretKey())
	if err != nil {
		return false, err
	}

	var info *model.UserInfo

	res, err := r.redis.Get(claims.Username).Result()
	if errors.Is(err, redis.Nil) {
		conditions := filter.MakeFilter(filter.Condition{
			Key:   model.UserNameFieldCode,
			Value: claims.Username,
		})

		user, errRep := r.userRepo.Get(ctx, conditions)
		if errRep != nil {
			return false, errRep
		}

		info = &user.Info
	}
	if err != nil {
		return false, err
	}

	if info == nil {
		err = json.Unmarshal([]byte(res), &info)
		if err != nil {
			return false, err
		}
	}

	if info.Role == model.ADMIN && claims.Role == model.ADMIN {
		return true, nil
	}

	res, err = r.redis.Get(endpoint).Result()
	if errors.Is(err, redis.Nil) {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	var roles []model.UserRole
	err = json.Unmarshal([]byte(res), &roles)
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role == info.Role {
			return true, nil
		}
	}

	return false, nil
}
