package model

import "time"

type Log struct {
	Action    string    `db:"action"`
	Content   string    `db:"content"`
	timestamp time.Time `db:"timestamp"`
}
