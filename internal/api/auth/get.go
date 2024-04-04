package authapi

import (
	"context"

	"github.com/a1exCross/auth/pkg/auth_v1"

	"github.com/pkg/errors"
)

// GetRefreshToken - возвращает новый refresh токен
func (i *Implementation) GetRefreshToken(ctx context.Context, req *auth_v1.GetRefreshTokenRequest) (*auth_v1.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, errors.Errorf("refresh token update error: %s", err)
	}

	return &auth_v1.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}

// GetAccessToken - возвращает новый access токен
func (i *Implementation) GetAccessToken(ctx context.Context, req *auth_v1.GetAccessTokenRequest) (*auth_v1.GetAccessTokenResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, errors.Errorf("access token get error: %s", err)
	}

	return &auth_v1.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
