package user

import (
	"fmt"
	_entities "group-project/dolan-planner/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreatUser(_entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"})
		assert.Nil(t, err)
		assert.Equal(t, "Usamah", data.Name)
	})

	t.Run("TestCreateUserErrorValidation1", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreatUser(_entities.User{Name: "Usamah", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"})
		assert.NotNil(t, err)
		assert.Equal(t, "Usamah", data.Name)
	})

	t.Run("TestCreateUserErrorValidation2", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreatUser(_entities.User{City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"})
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Name)
	})

	t.Run("TestCreateUserErrorValidation3", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreatUser(_entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Password: "usamah"})
		assert.NotNil(t, err)
		assert.Equal(t, "Usamah", data.Name)
	})

	t.Run("TestCreateUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.CreatUser(_entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.GetUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "Usamah", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.GetUser(1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{}, data)
		assert.Equal(t, 0, rows)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser((_entities.User{Name: "abdur"}), 1)
		assert.Nil(t, err)
		assert.Equal(t, "Usamah", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.UpdateUser((_entities.User{Name: "abdur"}), 1)
		assert.NotNil(t, err)
		assert.Equal(t, data, data)
		assert.Equal(t, 0, rows)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		rows, err := userUseCase.DeleteUser(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestDeleteUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		rows, err := userUseCase.DeleteUser(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
	})
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) CreatUser(createUser _entities.User) (_entities.User, error) {
	user := _entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}
	return user, nil
}

func (m mockUserRepository) GetUser(idToken int) (_entities.User, int, error) {
	return _entities.User{
		Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah",
	}, 1, nil
}

func (m mockUserRepository) UpdateUser(userUpdate _entities.User) (_entities.User, int, error) {
	user := _entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}
	return user, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) (int, error) {
	return 1, nil
}

// === mock error ===
type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) CreatUser(user _entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create user")
}

func (m mockUserRepositoryError) GetUser(idToken int) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error get user")
}

func (m mockUserRepositoryError) UpdateUser(userUpdate _entities.User) (_entities.User, int, error) {
	return _entities.User{Name: "usamah", City: "bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}, 0, fmt.Errorf("error update user")
}

func (m mockUserRepositoryError) DeleteUser(id int) (int, error) {
	return 0, fmt.Errorf("error delete user")
}
