package authservice

import (
	"github.com/a1exCross/auth/internal/config"
	"github.com/a1exCross/auth/internal/repository"
	"github.com/a1exCross/auth/internal/service"

	"github.com/a1exCross/common/pkg/storage"
)

type serv struct {
	redis     storage.Redis
	userRepo  repository.UserRepository
	jwtConfig config.JWTConfig
}

// NewService - создает экземпляр сервиса авторизации
func NewService(redis storage.Redis, repo repository.UserRepository, jwtConfig config.JWTConfig) service.AuthService {
	return &serv{
		redis:     redis,
		userRepo:  repo,
		jwtConfig: jwtConfig,
	}
}
