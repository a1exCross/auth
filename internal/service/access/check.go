package accessservice

import (
	"context"
	"errors"
	"strings"

	"github.com/a1exCross/auth/internal/model"

	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

func (s *serv) Check(ctx context.Context, roles ...model.UserRole) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	access, err := s.accessChecker.AccessCheck(ctx, accessToken, roles...)
	if err != nil {
		return err
	}

	if !access {
		return errors.New("access denied")
	}

	return nil
}