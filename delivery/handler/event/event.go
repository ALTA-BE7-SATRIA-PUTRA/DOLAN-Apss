package event

import (
	"group-project/dolan-planner/delivery/helper"
	_middlewares "group-project/dolan-planner/delivery/middlewares"
	"group-project/dolan-planner/entities"
	_eventUseCase "group-project/dolan-planner/usecase/event"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventUseCase _eventUseCase.EventUseCaseInterface
}

func NewEventHandler(eventUseCase _eventUseCase.EventUseCaseInterface) *EventHandler {
	return &EventHandler{
		eventUseCase: eventUseCase,
	}
}

func (eh *EventHandler) CreateEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var newEvent entities.Event
		err := c.Bind(&newEvent)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		_, error := eh.eventUseCase.CreatEvent(newEvent, uint(idToken))
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create event"))
	}
}
