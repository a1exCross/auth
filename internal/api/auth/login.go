package authapi

import (
	"context"
	"fmt"

	"github.com/a1exCross/auth/internal/api/auth/converter"
	"github.com/a1exCross/auth/pkg/auth_v1"
)

// Login - проводит аутентификацию пользователя
func (i *Implementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	refreshToken, err := i.authService.Login(ctx, converter.AuthProtoToAuthDTO(req))
	if err != nil {
		return nil, fmt.Errorf("authentification error: %s", err)
	}

	return &auth_v1.LoginResponse{
		RefreshToken: refreshToken,
	}, err
}
