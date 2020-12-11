package usecase_test

import (
	"go-store/domain"
	"go-store/repository/mocks"
	"go-store/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsecaseUser(t *testing.T) {
	mockUserRepo := &mocks.UserRepository{}
	mockUser := &domain.User{
		ID:        1,
		Username:  "TestUser",
		Password:  "12345678",
		CreatedAt: time.Now(),
	}

	mockUserRepo.On("User", mock.AnythingOfType("int64")).Return(mockUser, nil).Once()

	u := usecase.NewInteractor(mockUserRepo)

	user, err := u.User(mockUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseUsers(t *testing.T) {
	mockUserRepo := &mocks.UserRepository{}
	mockUser := &domain.User{
		ID:        1,
		Username:  "TestUser",
		Password:  "12345678",
		CreatedAt: time.Now(),
	}

	mockUserList := make([]*domain.User, 0)
	mockUserList = append(mockUserList, mockUser)

	mockUserRepo.On("Users", mock.Anything).Return(mockUserList, nil).Once()

	u := usecase.NewInteractor(mockUserRepo)

	users, err := u.Users()

	assert.NoError(t, err)
	assert.NotNil(t, users)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseCreateUser(t *testing.T) {
	mockUserRepo := &mocks.UserRepository{}
	mockUser := domain.User{
		ID:        1,
		Username:  "TestUser",
		Password:  "12345678",
		CreatedAt: time.Now(),
	}

	tempMockUser := mockUser

	mockUserRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(int64(1), nil).Once()

	u := usecase.NewInteractor(mockUserRepo)

	err := u.CreateUser(&tempMockUser)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.Username, tempMockUser.Username)
	mockUserRepo.AssertExpectations(t)
}

func TestUsecaseDeleteUser(t *testing.T) {
	mockUserRepo := &mocks.UserRepository{}
	mockUser := &domain.User{
		ID:        1,
		Username:  "TestUser",
		Password:  "12345678",
		CreatedAt: time.Now(),
	}

	mockUserRepo.On("DeleteUser", mock.AnythingOfType("int64")).Return(nil).Once()

	u := usecase.NewInteractor(mockUserRepo)

	err := u.DeleteUser(mockUser.ID)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}
