package userAPI

import (
	"context"
	"fmt"
	pbUser "github.com/a1exCross/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i Implementation) Delete(ctx context.Context, req *pbUser.DeleteRequest) (*empty.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	return &empty.Empty{}, nil
}
