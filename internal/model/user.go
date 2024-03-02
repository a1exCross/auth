package model

import (
	"database/sql"
	"time"
)

type UserRole int8

const (
	UNKNOWN UserRole = iota
	USER
	ADMIN
)

type User struct {
	ID        int64        `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Name  string   `db:"name"`
	Email string   `db:"email"`
	Role  UserRole `db:"role"`
}

type UserCreate struct {
	Info     UserInfo `db:""`
	Password string   `db:"password"`
}

type UserUpdate struct {
	ID   int64    `db:"id"`
	Info UserInfo `db:""`
}
