package interceptor

import (
	"context"
	"time"

	"github.com/a1exCross/auth/internal/logger"
	"github.com/a1exCross/auth/internal/metric"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// LoggingInterceptor - интерцептор для логирования действий и изменения метрик
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	metric.IncRequestCounter()

	timeStart := time.Now()

	res, err := handler(ctx, req)
	diffTime := time.Since(timeStart).Seconds()

	if err != nil {
		logger.Error("error at gRPC method call", zap.String("method", info.FullMethod), zap.Error(err))

		metric.IncResponseCounter("error", info.FullMethod)
		metric.HistogramResponseTimeObserve("error", diffTime)

		return nil, err
	}

	logger.Info("request validated successfully and done", zap.Any("req", req))

	metric.IncResponseCounter("success", info.FullMethod)
	metric.HistogramResponseTimeObserve("success", diffTime)

	return res, nil
}
