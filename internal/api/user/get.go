package userAPI

import (
	"context"
	"fmt"
	"github.com/a1exCross/auth/internal/api/user/converter"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"
)

func (i Implementation) Get(ctx context.Context, req *pbUser.GetRequest) (*pbUser.GetResponse, error) {
	res, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &pbUser.GetResponse{
		User: converter.UserToProto(res),
	}, err
}
