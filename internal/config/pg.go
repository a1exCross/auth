package config

import (
	"errors"
	"os"
)

const (
	pgDSNEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

// NewPGConfig - создание конфига для БД
func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(pgDSNEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("dsn not found in environments")
	}

	return pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg pgConfig) DSN() string {
	return cfg.dsn
}
