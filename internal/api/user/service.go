package userAPI

import (
	"github.com/a1exCross/auth/internal/service"
	userPb "github.com/a1exCross/auth/pkg/user_v1"
)

type Implementation struct {
	userPb.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(service service.UserService) *Implementation {
	return &Implementation{
		userService: service,
	}
}
