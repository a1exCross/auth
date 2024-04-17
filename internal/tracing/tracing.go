package tracing

import (
	"github.com/a1exCross/common/pkg/logger"

	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

// Init - иницилизирует конфиг Jaeger
func Init(serviceName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("failed to init tracing", zap.Error(err))
	}
}
