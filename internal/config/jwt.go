package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

const (
	jwtRefreshExpirationTime = "JWT_REFRESH_EXPIRATION_TIME"
	jwtRefreshSecretKey      = "JWT_REFRESH_SECRET_KEY" // #nosec G101
	jwtAccessExpirationTime  = "JWT_ACCESS_EXPIRATION_TIME"
	jwtAccessSecretKey       = "JWT_ACCESS_SECRET_KEY" // #nosec G101
)

type jwtConfig struct {
	refreshSecretKey      []byte
	refreshExpirationTime time.Duration
	accessSecretKey       []byte
	accessExpirationTime  time.Duration
}

// NewJWTConfig - создает конфиг JWT
func NewJWTConfig() (JWTConfig, error) {
	secretKeyRefresh := os.Getenv(jwtRefreshSecretKey)
	if len(secretKeyRefresh) == 0 {
		return nil, errors.New("jwt secret-key not found in environments")
	}

	expirationTimeRefresh := os.Getenv(jwtRefreshExpirationTime)
	if len(expirationTimeRefresh) == 0 {
		return nil, errors.New("jwt expiration token time not found in environments")
	}

	durationRefresh, err := strconv.Atoi(expirationTimeRefresh)
	if err != nil {
		return nil, errors.New("failed to convert string to int")
	}

	secretKeyAccess := os.Getenv(jwtAccessSecretKey)
	if len(secretKeyAccess) == 0 {
		return nil, errors.New("jwt secret-key not found in environments")
	}

	expirationTimeAccess := os.Getenv(jwtAccessExpirationTime)
	if len(expirationTimeAccess) == 0 {
		return nil, errors.New("jwt expiration token time not found in environments")
	}

	durationAccess, err := strconv.Atoi(expirationTimeAccess)
	if err != nil {
		return nil, errors.New("failed to convert string to int")
	}

	return jwtConfig{
		refreshExpirationTime: time.Duration(durationRefresh) * time.Minute,
		refreshSecretKey:      []byte(secretKeyRefresh),
		accessExpirationTime:  time.Duration(durationAccess) * time.Minute,
		accessSecretKey:       []byte(secretKeyAccess),
	}, nil
}

func (cfg jwtConfig) RefreshSecretKey() []byte {
	return cfg.refreshSecretKey
}

func (cfg jwtConfig) RefreshExpirationTime() time.Duration {
	return cfg.refreshExpirationTime
}

func (cfg jwtConfig) AccessSecretKey() []byte {
	return cfg.accessSecretKey
}

func (cfg jwtConfig) AccessExpirationTime() time.Duration {
	return cfg.accessExpirationTime
}
