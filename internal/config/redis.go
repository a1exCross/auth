package config

import (
	"errors"
	"net"
	"os"
)

const (
	redisHostEnvName = "REDIS_HOST"
	redisPortEnvName = "REDIS_PORT"
)

type redisConfig struct {
	host string
	port string
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

	return redisConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
