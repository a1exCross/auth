package config

import (
	"errors"
	"net"
	"os"
)

const (
	redisHostEnvName = "REDIS_HOST"
	redisPortEnvName = "REDIS_PORT"
	redisPassword    = "REDIS_PASSWORD"
)

type redisConfig struct {
	host     string
	port     string
	password string
}

// NewRedisConfig - создает конфиг Redis
func NewRedisConfig() (RedisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("redis host not found in environments")
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("redis port not found in environments")
	}

	password := os.Getenv(redisPassword)
	if len(password) == 0 {
		return nil, errors.New("redis password not found in environments")
	}

	return &redisConfig{
		host:     host,
		port:     port,
		password: password,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *redisConfig) Password() string {
	return cfg.password
}
