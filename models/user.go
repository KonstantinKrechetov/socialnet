package models

import "time"

type User struct {
	Id           int64     `db:"id"`
	Username     string    `db:"username"`
	FirstName    string    `db:"first_name"`
	SecondName   string    `db:"second_name"`
	Birthdate    time.Time `db:"birthdate"`
	Biography    string    `db:"biography"`
	City         string    `db:"city"`
	Password     string    `db:"password"`
	PasswordSalt string    `db:"password_salt"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
