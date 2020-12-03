package domain

import (
	"time"
)

var schema = `
CREATE TABLE user (
	id serial NOT NULL,
	username VARCHAR(25) NULL,
	password text,
	created_at date,
);`

type User struct {
	ID uint64
	Username string
	Password string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type UserRepository interface {
	User(id int64) (*User, error)
	Users() ([] *User, error)
	CreateUser(user *User) (int64, map[string]string)
	DeleteUser(id int64) error
}