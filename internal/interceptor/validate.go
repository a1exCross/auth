package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// Validator - интерфейс валидатора gRPC
type Validator interface {
	Validate() error
}

// ValidateInterceptor - валидатор запросов gRPC
func ValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if val, ok := req.(Validator); ok {
		err := val.Validate()
		if err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
