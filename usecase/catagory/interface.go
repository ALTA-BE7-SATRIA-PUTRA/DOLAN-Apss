package product

import (
	_entities "group-project/dolan-planner/entities"
)

type CatagoryUseCaseInterface interface {
	GetAllCatagory() ([]_entities.Catagory, error)
}
