package domain

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var schema = `
CREATE TABLE user (
	id serial NOT NULL,
	username VARCHAR(25) NULL,
	password text,
	created_at date,
);`

type User struct {
	ID        uint64
	Username  string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func NewUser(username, password string) (*User, error) {
	user := &User{
		Username:  strings.ToLower(username),
		CreatedAt: time.Now(),
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd
	err = user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() (err error) {
	if u.Username != " " && u.Password != " " {
		return err
	}
	return nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
