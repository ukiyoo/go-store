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
	user, err := interactor.repo.User(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (interactor *Interactor) Users() ([]*domain.User, error) {
	users, err := interactor.repo.Users()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (interactor *Interactor) CreateUser(user *domain.User) error {
	user, err := domain.NewUser(user.Username, user.Password)
	if err != nil {
		return err
	}
	_, err = interactor.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return err
}

func (interactor *Interactor) DeleteUser(id int64) error {
	err := interactor.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return err
}
