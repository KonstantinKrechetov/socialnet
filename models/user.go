package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `db:"id"`
	FirstName    string    `db:"first_name"`
	SecondName   string    `db:"second_name"`
	Birthdate    time.Time `db:"birthdate"`
	Biography    string    `db:"biography"`
	City         string    `db:"city"`
	PasswordHash string    `db:"password_hash"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
