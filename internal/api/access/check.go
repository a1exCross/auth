package accessapi

import (
	"context"
	"errors"
	"log"

	"github.com/a1exCross/auth/internal/api/access/converter"
	accesspb "github.com/a1exCross/auth/pkg/access_v1"

	"github.com/golang/protobuf/ptypes/empty"
)

// Check - проверяет доступность ручки для пользователя
func (i *Implementation) Check(ctx context.Context, req *accesspb.CheckRequest) (*empty.Empty, error) {
	log.Println(req.Role)

	err := i.service.Check(ctx, converter.RolesFromProto(req.GetRole())...)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return &empty.Empty{}, nil
}
