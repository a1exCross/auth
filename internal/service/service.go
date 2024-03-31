package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -o ./mocks/ -s ".go"

import (
	"context"

	"github.com/a1exCross/auth/internal/model"
)

// UserService - интерфейс, описывающий сервисный слой для работы с пользователями
type UserService interface {
	Create(context.Context, *model.UserCreate) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Delete(context.Context, int64) error
	Update(context.Context, *model.UserUpdate) error
}

// AuthService - сервис авторизации и аутентификации
type AuthService interface {
	Login(context.Context, model.LoginDTO) (string, error)
	GetRefreshToken(context.Context, string) (string, error)
	GetAccessToken(context.Context, string) (string, error)
}

// AccessService - сервис проверки доступов
type AccessService interface {
	Check(context.Context, ...model.UserRole) error
}
