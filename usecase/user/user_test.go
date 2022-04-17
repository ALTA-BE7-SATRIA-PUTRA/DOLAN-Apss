package user

import (
	"errors"
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
}

func TestGetUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.GetUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "Usamah", data.Name)
		assert.Equal(t, 1, rows)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("TestUpdateUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.UpdateUser((_entities.User{Name: "abdur"}), 1)
		assert.Nil(t, err)
		assert.Equal(t, "abdur", data.Name)
		assert.Equal(t, 1, rows)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestDeleteUser", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		rows, err := userUseCase.DeleteUser(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) CreatUser(createUser _entities.User) (_entities.User, error) {
	user := _entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}

	if createUser.Name == "" {
		return user, errors.New("name can't be empty")
	}
	if createUser.City == "" {
		return user, errors.New("city can't be empty")
	}
	if createUser.Email == "" {
		return user, errors.New("email can't be empty")
	}
	if createUser.Password == "" {
		return user, errors.New("password can't be empty")
	}

	return user, nil
}

func (m mockUserRepository) GetUser(idToken int) (_entities.User, int, error) {
	return _entities.User{
		Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah",
	}, 1, nil
}

func (m mockUserRepository) UpdateUser(userUpdate _entities.User, id int) (_entities.User, int, error) {
	user := _entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}

	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}
	if userUpdate.Email != "" {
		user.Email = userUpdate.Email
	}
	if userUpdate.Password != "" {
		user.Password = userUpdate.Password
	}
	if userUpdate.City != "" {
		user.City = userUpdate.City
	}
	if userUpdate.UrlImage != "" {
		user.UrlImage = userUpdate.UrlImage
	}

	return user, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) (int, error) {
	return 1, nil
}
