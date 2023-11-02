package models

import "time"

type Item struct {
	Id          int64     `db:"id,omitempty"`
	Name        string    `db:"name,omitempty"`
	Description string    `db:"description,omitempty"`
	CreatedAt   time.Time `db:"created_at"`
}
