package userAPI

import (
	"context"
	"fmt"
	"github.com/a1exCross/auth/internal/model"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i Implementation) Update(ctx context.Context, req *pbUser.UpdateRequest) (*empty.Empty, error) {
	err := i.userService.Update(ctx, &model.UserUpdate{
		Info: model.UserInfo{
			Name:  req.Name.Value,
			Email: req.Email.Value,
			Role:  model.UserRole(req.Role),
		},
		ID: req.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return &empty.Empty{}, nil
}
