package user

import (
	"errors"
	"group-project/dolan-planner/delivery/helper"
	_entities "group-project/dolan-planner/entities"
	_userRepository "group-project/dolan-planner/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) CreatUser(user _entities.User) (_entities.User, error) {
	password, _ := helper.HashPassword(user.Password)
	user.Password = password
	createUser, err := uuc.userRepository.CreatUser(user)

	if user.Name == "" {
		return user, errors.New("name can't be empty")
	}
	if user.City == "" {
		return user, errors.New("city can't be empty")
	}
	if user.Email == "" {
		return user, errors.New("email can't be empty")
	}
	if user.Password == "" {
		return user, errors.New("password can't be empty")
	}

	return createUser, err
}

func (uuc *UserUseCase) GetUser(idToken int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetUser(idToken)
	return user, rows, err
}

func (uuc *UserUseCase) UpdateUser(userUpdate _entities.User, id int) (_entities.User, int, error) {
	password, _ := helper.HashPassword(userUpdate.Password)
	userUpdate.Password = password
	user, rows, err := uuc.userRepository.GetUser(id)
	if err != nil {
		return user, 0, err
	}
	if rows == 0 {
		return user, 0, nil
	}
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

	updateUser, updateRows, updateErr := uuc.userRepository.UpdateUser(user)
	return updateUser, updateRows, updateErr
}

func (uuc *UserUseCase) DeleteUser(id int) (int, error) {
	rows, err := uuc.userRepository.DeleteUser(id)
	return rows, err
}
