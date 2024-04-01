package accessapi

import (
	"context"
	"errors"

	accesspb "github.com/a1exCross/auth/pkg/access_v1"

	"github.com/golang/protobuf/ptypes/empty"
)

// Check - проверяет доступность ручки для пользователя
func (i *Implementation) Check(ctx context.Context, req *accesspb.CheckRequest) (*empty.Empty, error) {
	err := i.service.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, errors.New("access denied")
	}

	return &empty.Empty{}, nil
}
