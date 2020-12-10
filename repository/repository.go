package repository

import "go-store/domain"

type UserRepository interface {
	User(id int64) (*domain.User, error)
	Users() ([]*domain.User, error)
	CreateUser(user *domain.User) (int64, error)
	DeleteUser(id int64) error
}
