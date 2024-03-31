package config

import (
	"time"

	"github.com/joho/godotenv"
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
