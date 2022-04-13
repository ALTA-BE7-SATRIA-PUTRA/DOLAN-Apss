package attendees

import (
	"group-project/dolan-planner/delivery/helper"
	_middlewares "group-project/dolan-planner/delivery/middlewares"
	_attendeesUseCase "group-project/dolan-planner/usecase/attendees"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AttendeesHandler struct {
	attendeesUseCase _attendeesUseCase.AttendeesUseCaseInterface
}

func NewAttendeesHandler(attendeesUseCase _attendeesUseCase.AttendeesUseCaseInterface) *AttendeesHandler {
	return &AttendeesHandler{
		attendeesUseCase: attendeesUseCase,
	}
}

func (uh *AttendeesHandler) PostAttendeesHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		idEvent, errorconv := strconv.Atoi(idStr)
		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be number id event")
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		_, idErr, _ := uh.attendeesUseCase.PostAttendees(uint(idEvent), uint(idToken))
		if idErr == 1 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail to read event"))
		}

		if idErr == 2 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("quota is full"))
		}

		if idErr == 3 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("you have joined"))
		}

		if idErr == 4 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail to read attendees"))
		}
		if idErr == 6 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("event not found"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succses join to event"))

	}
}

func (uh *AttendeesHandler) GetAttendeesHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		idEvent, errorconv := strconv.Atoi(idStr)
		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be number id event")
		}

		attendees, err := uh.attendeesUseCase.GetAttendees(uint(idEvent))
		errString := err.Error()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errString))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to get attendees", attendees))
	}
}
