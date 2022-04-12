package catagory

import (
	"group-project/dolan-planner/delivery/helper"
	"group-project/dolan-planner/entities"
	_catagoryUseCase "group-project/dolan-planner/usecase/catagory"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CatagoryHandler struct {
	catagoryUseCase _catagoryUseCase.CatagoryUseCaseInterface
}

func NewCatagoryHandler(c _catagoryUseCase.CatagoryUseCaseInterface) CatagoryHandler {
	return CatagoryHandler{
		catagoryUseCase: c,
	}
}

func (uh *CatagoryHandler) CreateCatagoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newCatagory entities.Catagory
		err := c.Bind(&newCatagory)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		_, error := uh.catagoryUseCase.CreateCatagory(newCatagory)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create catagory"))
	}
}

func (uh *CatagoryHandler) GetAllCatagoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		catagory, err := uh.catagoryUseCase.GetAllCatagory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseCatagories := []map[string]interface{}{}
		for i := 0; i < len(catagory); i++ {
			response := map[string]interface{}{
				"id":            catagory[i].ID,
				"catagory_name": catagory[i].CatagoryName,
			}
			responseCatagories = append(responseCatagories, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all catagories", responseCatagories))
	}
}
