package service

import (
	"context"
	"github.com/a1exCross/auth/internal/model"
)

type UserService interface {
	Create(context.Context, *model.UserCreate) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Delete(context.Context, int64) error
	Update(context.Context, *model.UserUpdate) error
}
