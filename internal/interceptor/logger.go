package interceptor

import (
	"context"

	"github.com/a1exCross/common/pkg/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// LoggingInterceptor - интерцептор для логирования действий и изменения метрик
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if err != nil {
		logger.Error("error at gRPC method call", zap.String("method", info.FullMethod), zap.Error(err))

		return nil, err
	}

	logger.Info("request validated successfully and done", zap.Any("req", req))

	return res, nil
}
