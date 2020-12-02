package repository

import "go-store/domain"

type UserRepository interface {
	User(id uint64) (*domain.User, error)
	Users() ([] *domain.User, error)
	CreateUser(user *domain.User) (int64, map[string]string)
	DeleteUser(id uint64) error
}