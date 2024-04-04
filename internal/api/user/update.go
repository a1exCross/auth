package userapi

import (
	"context"

	"github.com/a1exCross/auth/internal/api/user/converter"
	"github.com/a1exCross/auth/internal/model"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

// Update принимает и обрабатывает запрос на обновление пользователя
func (i *Implementation) Update(ctx context.Context, req *pbUser.UpdateRequest) (*empty.Empty, error) {
	err := i.userService.Update(ctx, &model.UserUpdate{
		Info: converter.ProtoToUserInfoUpdate(req.GetInfo()),
		ID:   req.Id,
	})
	if err != nil {
		return nil, errors.Errorf("failed to update user: %v", err)
	}

	return &empty.Empty{}, nil
}
