package product

import (
	"fmt"
	_entities "group-project/dolan-planner/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCatagory(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		categoryUseCase := NewCatagoryUseCase(mockCategoryRepository{})
		data, err := categoryUseCase.GetAllCatagory()
		assert.Nil(t, err)
		assert.Equal(t, "category 1", data[0].CatagoryName)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		categoryUseCase := NewCatagoryUseCase(mockCategoryRepositoryError{})
		data, err := categoryUseCase.GetAllCatagory()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestCreateCatagory(t *testing.T) {
	t.Run("TestCreateCatagorySuccess", func(t *testing.T) {
		catagoryUseCase := NewCatagoryUseCase(mockCategoryRepository{})
		catagory, err := catagoryUseCase.CreateCatagory(_entities.Catagory{CatagoryName: "sport"})
		assert.Nil(t, err)
		assert.Equal(t, "sport", catagory.CatagoryName)
	})

	t.Run("TestCreateCatagoryError", func(t *testing.T) {
		catagoryUseCase := NewCatagoryUseCase(mockCategoryRepositoryError{})
		catagory, err := catagoryUseCase.CreateCatagory(_entities.Catagory{CatagoryName: "sport"})
		assert.NotNil(t, err)
		assert.Equal(t, "", catagory.CatagoryName)
	})
}

// === mock success ===
type mockCategoryRepository struct{}

func (m mockCategoryRepository) GetAllCatagory() ([]_entities.Catagory, error) {
	return []_entities.Catagory{
		{CatagoryName: "category 1"},
	}, nil
}

func (m mockCategoryRepository) CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error) {
	return _entities.Catagory{CatagoryName: "sport"}, nil
}

// === mock error ===

type mockCategoryRepositoryError struct{}

func (m mockCategoryRepositoryError) GetAllCatagory() ([]_entities.Catagory, error) {
	return nil, fmt.Errorf("error get all data user")
}

func (m mockCategoryRepositoryError) CreateCatagory(catagory _entities.Catagory) (_entities.Catagory, error) {
	return _entities.Catagory{}, fmt.Errorf("error create catagory")
}
