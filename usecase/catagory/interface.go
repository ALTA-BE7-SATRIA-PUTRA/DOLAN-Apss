package product

import (
	_entities "group-project/dolan-planner/entities"
)

type CatagoryUseCaseInterface interface {
	CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error)
	GetAllCatagory() ([]_entities.Catagory, error)
}
