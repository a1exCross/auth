package config

import (
	"time"

	"github.com/a1exCross/auth/internal/model"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Load - парсит .env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// RedisConfig - конфиг Redis
type RedisConfig interface {
	Address() string
	Password() string
	RoutesAccesses() map[string][]model.UserRole
}

// SwaggerConfig - конфиг swagger
type SwaggerConfig interface {
	Address() string
}

// HTTPConfig - конфиг http
type HTTPConfig interface {
	Address() string
}

// GRPCConfig - конфиг gRPC
type GRPCConfig interface {
	Address() string
}

// PGConfig - конфиг Postgres
type PGConfig interface {
	DSN() string
}

// JWTConfig - конфиг JWT
type JWTConfig interface {
	RefreshSecretKey() []byte
	RefreshExpirationTime() time.Duration
	AccessSecretKey() []byte
	AccessExpirationTime() time.Duration
}

// LoggerConfig - конфиг логгера
type LoggerConfig interface {
	getCore() zapcore.Core
	getLogLevel() zap.AtomicLevel
}

// PrometheusConfig - конфиг prometheus
type PrometheusConfig interface {
	Address() string
}
