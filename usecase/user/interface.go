package user

import (
	_entities "group-project/dolan-planner/entities"
)

type UserUseCaseInterface interface {
	CreatUser(user _entities.User) (_entities.User, error)
	GetUser(idToken int) (_entities.User, int, error)
	UpdateUser(userUpdate _entities.User, id int) (_entities.User, int, error)
	DeleteUser(id int) (int, error)
}
