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

func (u *User) Validate() (err error)  {
	if u.Username != " " && u.Password != " " {
		return err
	}
	return nil
}