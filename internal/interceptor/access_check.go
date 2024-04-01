package interceptor

import (
	"context"
	"errors"
	"strings"

	"github.com/a1exCross/auth/internal/service"

	"google.golang.org/grpc"
)

// AccessChecker - верификатор доступа
type AccessChecker interface {
	AccessCheck(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}

type checker struct {
	authClient service.AccessService
}

// NewAccessChecker - создание верификатора доступа
func NewAccessChecker(auth service.AccessService) AccessChecker {
	return &checker{
		authClient: auth,
	}
}

// AccessCheck - проверка доступа к ручке
func (c *checker) AccessCheck(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if strings.Contains(info.FullMethod, "access_v1") || strings.Contains(info.FullMethod, "auth_v1") ||
		info.FullMethod == "/user_v1.UserV1/Create" {
		return handler(ctx, req)
	}

	err := c.authClient.Check(ctx, info.FullMethod)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return handler(ctx, req)
}
