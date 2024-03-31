package utils

import (
	"time"

	"github.com/a1exCross/auth/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// GenerateToken - генерирует JWT токен
func GenerateToken(info model.UserInfo, secretKey []byte, duration time.Duration) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Username: info.Username,
		Role:     info.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

// VerifyToken - верифицирует JWT токен
func VerifyToken(tokenHash string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenHash,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("wrong type signing method")
			}

			return secretKey, nil
		})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
