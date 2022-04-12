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

func (eh *EventHandler) GetAllEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := eh.eventUseCase.GetAllEvent()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseEvents := []map[string]interface{}{}
		for i := 0; i < len(events); i++ {
			response := map[string]interface{}{
				"id":         events[i].ID,
				"name_event": events[i].NameEvent,
				"date":       events[i].Date,
				"hosted_by":  events[i].HostedBy,
				"location":   events[i].Location,
			}
			responseEvents = append(responseEvents, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all events", responseEvents))
	}
}
