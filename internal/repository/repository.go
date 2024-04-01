package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -o ./mocks/ -s ".go"

import (
	"context"

	"github.com/a1exCross/auth/internal/model"

	"github.com/a1exCross/common/pkg/filter"
)

// UserRepository - описывает методы репозитория пользователей
type UserRepository interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	Get(ctx context.Context, filter filter.Filter) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user *model.UserUpdate) error
}

// LogsRepository - описывает методы репозитория логов
type LogsRepository interface {
	Create(ctx context.Context, log model.Log) (int64, error)
	Get(ctx context.Context, id int64) (model.Log, error)
}
