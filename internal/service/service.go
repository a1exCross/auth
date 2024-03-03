package service

import (
	"github.com/a1exCross/auth/internal/model"

	"context"
)

// UserService - интерфейс, описывающий сервисный слой для работы с пользователями
type UserService interface {
	Create(context.Context, *model.UserCreate) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Delete(context.Context, int64) error
	Update(context.Context, *model.UserUpdate) error
}
