package repository

import (
	"context"

	"github.com/a1exCross/auth/internal/model"
)

// UserRepository - описывает методы репозитория пользователей
type UserRepository interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user *model.UserUpdate) error
}

// LogsRepository - описывает методы репозитория логов
type LogsRepository interface {
	Create(ctx context.Context, log model.Log) (int64, error)
	Get(ctx context.Context, id int64) (model.Log, error)
}
