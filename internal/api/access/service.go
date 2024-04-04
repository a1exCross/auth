package accessapi

import (
	"github.com/a1exCross/auth/internal/service"
	accesspb "github.com/a1exCross/auth/pkg/access_v1"
)

// Implementation - структура, описывающая имплементацию gRPC сервера
type Implementation struct {
	accesspb.UnimplementedAccessV1Server
	service service.AccessService
}

// NewImplementation - Создает новую имплементацию для gRPC сервера
func NewImplementation(serv service.AccessService) *Implementation {
	return &Implementation{
		service: serv,
	}
}
