package userapi

import (
	"context"

	pbUser "github.com/a1exCross/auth/pkg/user_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

// Delete принимает и обрабатывает запрос на удаление пользователя
func (i *Implementation) Delete(ctx context.Context, req *pbUser.DeleteRequest) (*empty.Empty, error) {
	err := i.accessService.Check(ctx, "/user_v1.UserV1/Delete")
	if err != nil {
		return nil, err
	}

	err = i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, errors.Errorf("failed to delete user: %v", err)
	}

	return &empty.Empty{}, nil
}
