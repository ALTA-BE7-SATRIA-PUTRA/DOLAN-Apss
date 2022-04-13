package catagory

import (
	_entities "group-project/dolan-planner/entities"

	"gorm.io/gorm"
)

type CatagoryRepository struct {
	DB *gorm.DB
}

func NewCatagoryRepository(db *gorm.DB) *CatagoryRepository {
	return &CatagoryRepository{
		DB: db,
	}
}

func (cr *CatagoryRepository) CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error) {
	tx := cr.DB.Save(&catagory)
	if tx.Error != nil {
		return catagory, tx.Error
	}
	return catagory, nil
}

func (cr *CatagoryRepository) GetAllCatagory() ([]_entities.Catagory, error) {
	var catagory []_entities.Catagory
	tx := cr.DB.Find(&catagory)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return catagory, nil
}
