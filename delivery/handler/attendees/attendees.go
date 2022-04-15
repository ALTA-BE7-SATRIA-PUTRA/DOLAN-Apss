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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("The expected param must be number type integer"))
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		_, err := uh.attendeesUseCase.PostAttendees(uint(idEvent), uint(idToken))
		errString := err.Error()
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errString))
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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("The expected param must be number type integer"))
		}

		attendees, rows, err := uh.attendeesUseCase.GetAttendees(uint(idEvent))
		if rows == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("attendees not found"))
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get attendees"))
		}

		responseAttendees := []map[string]interface{}{}
		for i := 0; i < len(attendees); i++ {
			response := map[string]interface{}{
				"id":       attendees[i].ID,
				"event_id": attendees[i].EventId,
				"user_id":  attendees[i].UserId,
				"user": map[string]interface{}{
					"name":      attendees[i].User.Name,
					"url_image": attendees[i].User.UrlImage},
			}
			responseAttendees = append(responseAttendees, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to get attendees", responseAttendees))
	}
}

func (ah *AttendeesHandler) DeleteAttendeesHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		idEvent, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("The expected param must be number type integer"))
		}

		//jika id sama dan tidak ada error
		rows, err := ah.attendeesUseCase.DeleteAttendees(uint(idToken), uint(idEvent))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("successfully left the event"))
	}
}

func (uh *AttendeesHandler) GetAttendeesUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		attendees, rows, err := uh.attendeesUseCase.GetAttendeesUser(uint(idToken))
		if rows == 0 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("data not found"))
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		responseAttendees := []map[string]interface{}{}
		for i := 0; i < len(attendees); i++ {
			response := map[string]interface{}{
				"id":       attendees[i].ID,
				"event_id": attendees[i].EventId,
				"user_id":  attendees[i].UserId,
				"event": map[string]interface{}{
					"name_event": attendees[i].Event.NameEvent,
					"date":       attendees[i].Event.Date,
					"hosted_by":  attendees[i].Event.HostedBy,
					"location":   attendees[i].Event.Location,
					"url_image":  attendees[i].Event.UrlImage},
			}
			responseAttendees = append(responseAttendees, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to get attendees", responseAttendees))
	}
}
