package catagory

import (
	_entities "group-project/dolan-planner/entities"
)

type CatagoryRepositoryInterface interface {
	CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error)
	GetAllCatagory() ([]_entities.Catagory, error)
}
