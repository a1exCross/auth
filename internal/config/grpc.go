package config

import (
	"errors"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig - создает конфиг gRPC
func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found in environments")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc port not found in environments")
	}

	return grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
