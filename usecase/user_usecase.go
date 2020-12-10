package usecase

import (
	"go-store/domain"
	"go-store/repository"
)

type Interactor struct {
	repo repository.UserRepository
}

func NewInteractor(u repository.UserRepository) *Interactor {
	return &Interactor{repo: u}
}

func (interactor *Interactor) User(id int64) (*domain.User, error) {
	panic("Implement me")
}

func (interactor *Interactor) CreateUser(user *domain.User) error {
	u, err := domain.NewUser(user.Username, user.Password)
	if err != nil {
		return err
	}
	_, err = interactor.repo.CreateUser(u)
	if err != nil {
		return err
	}
	return err
}

func (interactor *Interactor) DeleteUser(id int64) (*domain.User, error) {
	panic("Implement me")
}
