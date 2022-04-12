package catagory

import (
	_entities "group-project/dolan-planner/entities"
)

type CatagoryRepositoryInterface interface {
	GetAllCatagory() ([]_entities.Catagory, error)
}
