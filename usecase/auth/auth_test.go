package auth

import (
	"fmt"
	_entities "group-project/dolan-planner/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepository{})
		token, err := authUseCase.Login("usamah@gmail.com", "usamaha")
		assert.Nil(t, err)
		assert.Equal(t, "ini token", token)
	})

	t.Run("TestLoginError", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepositoryError{})
		token, err := authUseCase.Login("usamah@gmail.com", "usamah")
		assert.NotNil(t, err)
		assert.Equal(t, "password incorrect", token)
	})
}

// === mock success ===
type mockAuthRepository struct{}

func (m mockAuthRepository) Login(email string, password string) (string, error) {
	user := _entities.User{Name: "usamah", Email: "usamah@gmail.com", Password: "usamaha"}

	if user.Password != password {
		return "password incorrect", fmt.Errorf("password incorrect")
	}
	return "ini token", nil
}

// === mock error ===
type mockAuthRepositoryError struct{}

func (m mockAuthRepositoryError) Login(email string, password string) (string, error) {
	user := _entities.User{Name: "usamah", Email: "usamah@gmail.com", Password: "usamaha"}

	// jika password yang di input tidak sesuai
	if user.Password != password {
		return "password incorrect", fmt.Errorf("password incorrect")
	}
	return "ini token", nil
}
