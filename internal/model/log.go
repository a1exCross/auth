package model

import "time"

// Log - структура, описывающая лог в БД
type Log struct {
	Action    string    `db:"action"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
