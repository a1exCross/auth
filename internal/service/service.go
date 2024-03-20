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
