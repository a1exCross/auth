package config

import (
	"errors"
	"net"
	"os"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

type swaggerConfig struct {
	host string
	port string
}

// NewSwaggerConfig - создает конфиг swagger
func NewSwaggerConfig() (SwaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found in environments")
	}

	port := os.Getenv(swaggerPortEnvName)
	if len(host) == 0 {
		return nil, errors.New("swagger port not found in environments")
	}

	return swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
