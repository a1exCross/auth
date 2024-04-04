package userapi

import (
	"context"

	"github.com/a1exCross/auth/internal/api/user/converter"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"

	"github.com/pkg/errors"
)

// Get принимает и обрабатывает запрос на получение пользователя
func (i *Implementation) Get(ctx context.Context, req *pbUser.GetRequest) (*pbUser.GetResponse, error) {
	res, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, errors.Errorf("failed to get user: %v", err)
	}

	return &pbUser.GetResponse{
		User: converter.UserToProto(res),
	}, err
}
