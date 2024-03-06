package model

import (
	"database/sql"
	"time"
)

// UserRole - Описывает роль пользователя от 0 до 255
type UserRole int8

// Типы ролей пользователя
const (
	UNKNOWN UserRole = iota
	USER
	ADMIN
)

// User - структура, оисывающая пользователя в БД
type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// UserInfo - структура, оисывающая информацию о пользователе в БД
type UserInfo struct {
	Name  string   `db:"name"`
	Email string   `db:"email"`
	Role  UserRole `db:"role"`
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
