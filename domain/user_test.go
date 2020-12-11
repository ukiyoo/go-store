package domain_test

import (
	"go-store/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := domain.NewUser("TestUser", "12345678")
	assert.NoError(t, err)
	assert.Equal(t, user.Username, "testuser")
	assert.NotEqual(t, user.Password, "12345678")
}

func TestValidatePassword(t *testing.T) {
	user, _ := domain.NewUser("TestUser", "password")
	err := user.ValidatePassword("password")
	assert.NoError(t, err)
	err = user.ValidatePassword("wrong_pass")
	assert.Error(t, err)
}

func TestValidate(t *testing.T) {
	user := domain.User{
		ID:        1,
		Username:  "TestUser",
		Password:  "12345678",
		CreatedAt: time.Now(),
	}
	t.Run("success-validate", func(t *testing.T) {
		err := user.Validate()
		assert.NoError(t, err)
		assert.Equal(t, user.Username, "TestUser")
	})
	t.Run("error-validate", func(t *testing.T) {
		user.Username = " "
		err := user.Validate()
		assert.Error(t, err)
		assert.Equal(t, user.Username, " ")
	})
}
