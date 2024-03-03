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

// GRPCConfig - конфиг gRPC
type GRPCConfig interface {
	Address() string
}

// PGConfig - конфиг Postgres
type PGConfig interface {
	DSN() string
}
