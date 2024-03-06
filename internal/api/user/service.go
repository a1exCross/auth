package userapi

import (
	"github.com/a1exCross/auth/internal/service"
	userPb "github.com/a1exCross/auth/pkg/user_v1"
)

// Implementation - структура, описывающая имплементацию gRPC сервера
type Implementation struct {
	userPb.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation - создает новую имплементацию для gRPC сервера
func NewImplementation(service service.UserService) *Implementation {
	return &Implementation{
		userService: service,
	}
}
