package user

import (
	_entities "group-project/dolan-planner/entities"
)

type UserRepositoryInterface interface {
	CreatUser(user _entities.User) (_entities.User, error)
	GetUser(idToken int) (_entities.User, int, error)
	UpdateUser(userUpdate _entities.User) (_entities.User, int, error)
	DeleteUser(id int) (int, error)
}
