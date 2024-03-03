package userapi

import (
	"github.com/a1exCross/auth/internal/api/user/converter"
	"github.com/a1exCross/auth/internal/model"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"

	"context"
	"fmt"
)

// Create принимает и обрабатывает запрос на создание пользователя
func (i Implementation) Create(ctx context.Context, req *pbUser.CreateRequest) (*pbUser.CreateResponse, error) {
	if req.Pass.Password != req.Pass.PasswordConfirm {
		return nil, fmt.Errorf("passwords mismatch")
	}

	res, err := i.userService.Create(ctx, &model.UserCreate{
		Info:     converter.ProtoToUserInfo(req.GetInfo()),
		Password: req.Pass.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &pbUser.CreateResponse{
		Id: res,
	}, nil
}
