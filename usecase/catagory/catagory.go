package product

import (
	_entities "group-project/dolan-planner/entities"
	_catagoryRepository "group-project/dolan-planner/repository/catagory"
)

type CatagoryUseCase struct {
	catagoryRepository _catagoryRepository.CatagoryRepositoryInterface
}

func NewCatagoryUseCase(catagoryRepo _catagoryRepository.CatagoryRepositoryInterface) CatagoryUseCaseInterface {
	return &CatagoryUseCase{
		catagoryRepository: catagoryRepo,
	}
}

func (cuc *CatagoryUseCase) CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error) {
	createCatagory, err := cuc.catagoryRepository.CreateCatagory(catagory)
	return createCatagory, err
}

func (cuc *CatagoryUseCase) GetAllCatagory() ([]_entities.Catagory, error) {
	catagory, err := cuc.catagoryRepository.GetAllCatagory()
	return catagory, err
}
