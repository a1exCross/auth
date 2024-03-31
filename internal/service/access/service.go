package accessservice

import (
	"github.com/a1exCross/auth/internal/config"
	"github.com/a1exCross/auth/internal/service"
	"github.com/a1exCross/auth/internal/utils"
)

type serv struct {
	jwtConfig     config.JWTConfig
	accessChecker utils.AccessChecker
}

// NewService - создает новый экземпляр сервиса проверки доступа
func NewService(jwtConfig config.JWTConfig, checker utils.AccessChecker) service.AccessService {
	return &serv{
		jwtConfig:     jwtConfig,
		accessChecker: checker,
	}
}
