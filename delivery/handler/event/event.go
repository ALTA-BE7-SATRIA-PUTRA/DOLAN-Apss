package event

import (
	"group-project/dolan-planner/delivery/helper"
	_middlewares "group-project/dolan-planner/delivery/middlewares"
	"group-project/dolan-planner/entities"
	_eventUseCase "group-project/dolan-planner/usecase/event"
	"net/http"
	"strconv"

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

		//proses binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err_binding_image.Error()))
		}
		// check file extension
		_, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err_check_extension.Error()))
		}
		// check file size
		err_check_size := helper.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err_check_size.Error()))
		}
		fileName := "event" + "_" + strconv.Itoa(idToken) + "_" + strconv.Itoa(int(newEvent.MaxParticipants))

		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := helper.UploadImage("imageEvent", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err_upload_photo.Error()))
		}

		newEvent.UrlImage = theUrl

		_, error := eh.eventUseCase.CreatEvent(newEvent, uint(idToken))
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(error.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create event"))
	}
}

func (eh *EventHandler) GetAllEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		filters := map[string]string{}
		if c.QueryParam("name_event") != "" {
			filters["name_event"] = c.QueryParam("name_event")
		}
		if c.QueryParam("location") != "" {
			filters["location"] = c.QueryParam("location")
		}

		events, err := eh.eventUseCase.GetAllEvent(filters)
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
				"url_image":  events[i].UrlImage,
			}
			responseEvents = append(responseEvents, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all events", responseEvents))
	}
}

func (eh *EventHandler) GetEventByIdHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		event, rows, err := eh.eventUseCase.GetEventById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		attendees := []map[string]interface{}{}
		for i := 0; i < len(event.Attendees); i++ {
			response := map[string]interface{}{
				"id": event.Attendees[i].ID,
				"user": map[string]interface{}{
					"user_id":   event.Attendees[i].User.ID,
					"name":      event.Attendees[i].User.Name,
					"url_image": event.Attendees[i].User.UrlImage},
			}
			attendees = append(attendees, response)
		}

		comment := []map[string]interface{}{}
		for i := 0; i < len(event.Comment); i++ {
			response := map[string]interface{}{
				"id":         event.Comment[i].ID,
				"created_at": event.Comment[i].CreatedAt,
				"comment":    event.Comment[i].Comment,
				"user": map[string]interface{}{
					"user_id": event.Comment[i].User.ID,
					"name":    event.Comment[i].User.Name},
			}
			comment = append(comment, response)
		}

		responseEvent := map[string]interface{}{
			"id":                 event.ID,
			"catagory_id":        event.ID,
			"name_event":         event.NameEvent,
			"hosted_by":          event.HostedBy,
			"max_participants":   event.MaxParticipants,
			"total_participants": event.TotalParticipants,
			"date":               event.Date,
			"location":           event.Location,
			"detail_event":       event.DetailEvent,
			"url_image":          event.UrlImage,
			"attendees":          attendees,
			"comment":            comment,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get event by id", responseEvent))
	}
}

func (eh *EventHandler) UpdateEventHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		var updateEvent entities.Event
		errBind := c.Bind(&updateEvent)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		event, rows, err := eh.eventUseCase.UpdateEvent(updateEvent, id, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseEvent := map[string]interface{}{
			"id":                 event.ID,
			"catagory_id":        event.ID,
			"name_event":         event.NameEvent,
			"hosted_by":          event.HostedBy,
			"max_participants":   event.MaxParticipants,
			"total_participants": event.TotalParticipants,
			"date":               event.Date,
			"location":           event.Location,
			"detail_event":       event.DetailEvent,
			"url_image":          event.UrlImage,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update event", responseEvent))
	}
}

func (eh *EventHandler) DeleteEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		// check apakah id dari token sama dengan id dari parm
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		//jika id sama dan tidak ada error
		rows, err := eh.eventUseCase.DeleteEvent(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success deleted event"))
	}
}
