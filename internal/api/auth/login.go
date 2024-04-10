package authapi

import (
	"context"

	"github.com/a1exCross/auth/internal/api/auth/converter"
	"github.com/a1exCross/auth/internal/logger"
	"github.com/a1exCross/auth/pkg/auth_v1"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Login - проводит аутентификацию пользователя
func (i *Implementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	refreshToken, err := i.authService.Login(ctx, converter.AuthProtoToAuthDTO(req))
	if err != nil {
		return nil, errors.Errorf("authentification error: %s", err)
	}

	logger.Info("user logged", zap.String("username", req.GetUsername()))

	return &auth_v1.LoginResponse{
		RefreshToken: refreshToken,
	}, err
}
