package config

import "github.com/joho/godotenv"

// Load - парсит .env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
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
