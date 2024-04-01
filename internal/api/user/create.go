package userapi

import (
	"context"

	"github.com/a1exCross/auth/internal/api/user/converter"
	"github.com/a1exCross/auth/internal/model"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"

	"github.com/pkg/errors"
)

// Create принимает и обрабатывает запрос на создание пользователя
func (i *Implementation) Create(ctx context.Context, req *pbUser.CreateRequest) (*pbUser.CreateResponse, error) {
	if req.Pass.Password != req.Pass.PasswordConfirm {
		return nil, errors.Errorf("passwords mismatch")
	}

	res, err := i.userService.Create(ctx, &model.UserCreate{
		Info:     converter.ProtoToUserInfo(req.GetInfo()),
		Password: req.Pass.Password,
	})
	if err != nil {
		return nil, errors.Errorf("failed to create user: %v", err)
	}

	return &pbUser.CreateResponse{
		Id: res,
	}, nil
}
