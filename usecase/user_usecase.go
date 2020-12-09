package usecase

import (
	"go-store/domain"
)

type Interactor struct {
	userInteractor UserUsecase
}

func NewInteractor(u UserUsecase) *Interactor {
	return &Interactor{userInteractor: u}
}

func (interactor *Interactor) User(id int64) (*domain.User, error) {
	panic("Implement me")
}