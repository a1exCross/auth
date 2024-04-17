package accessapi

import (
	"context"
	"errors"

	accesspb "github.com/a1exCross/auth/pkg/access_v1"

	"github.com/a1exCross/common/pkg/logger"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
)

// Check - проверяет доступность ручки для пользователя
func (i *Implementation) Check(ctx context.Context, req *accesspb.CheckRequest) (*empty.Empty, error) {
	logger.Info("Checking access request", zap.String("endpoint", req.EndpointAddress))

	err := i.service.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, errors.New("access denied")
	}

	return &empty.Empty{}, nil
}
