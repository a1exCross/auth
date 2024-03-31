package model

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserRole - Описывает роль пользователя от 0 до 255
type UserRole int8

// Типы ролей пользователя
const (
	UNKNOWN UserRole = iota
	USER
	ADMIN
)

// User - структура, описывающая пользователя в БД
type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Password  string       `db:"password"`
}

// UserInfo - структура, описывающая информацию о пользователе в БД
type UserInfo struct {
	Username string   `db:"username"`
	Name     string   `db:"name"`
	Email    string   `db:"email"`
	Role     UserRole `db:"role"`
}

// UserCreate - DTO для создания пользователя
type UserCreate struct {
	Info     UserInfo `db:""`
	Password string   `db:"password"`
}

// UserUpdate - DTO для обновления пользователя
type UserUpdate struct {
	ID   int64    `db:"id"`
	Info UserInfo `db:""`
}

// UserClaims - параметры для JWT токена
type UserClaims struct {
	jwt.StandardClaims
	Username string
	Role     UserRole
}
