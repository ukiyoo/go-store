package usecase

import "go-store/domain"

type UserUsecase interface {
	User(id int64) (*domain.User, error)
	Users() ([]*domain.User, error)
	CreateUser(user *domain.User) error
	DeleteUser(id int64) error
}
