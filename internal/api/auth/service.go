package authapi

import (
	"github.com/a1exCross/auth/internal/service"
	authPb "github.com/a1exCross/auth/pkg/auth_v1"
)

// Implementation - структура, описывающая имплементацию gRPC сервера
type Implementation struct {
	authPb.UnimplementedAuthV1Server
	authService service.AuthService
}

// NewImplementation - создает новую имплементацию для gRPC сервера
func NewImplementation(service service.AuthService) *Implementation {
	return &Implementation{
		authService: service,
	}
}
