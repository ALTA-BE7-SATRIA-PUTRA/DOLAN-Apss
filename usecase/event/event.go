package event

import (
	"errors"
	_entities "group-project/dolan-planner/entities"
	_eventRepository "group-project/dolan-planner/repository/event"
	_userRepository "group-project/dolan-planner/repository/user"
)

type EventUseCase struct {
	eventRepository _eventRepository.EventRepositoryInterface
	userRepository  _userRepository.UserRepositoryInterface
}

func NewEventUseCase(eventRepo _eventRepository.EventRepositoryInterface, userRepo _userRepository.UserRepositoryInterface) EventUseCaseInterface {
	return &EventUseCase{
		eventRepository: eventRepo,
		userRepository:  userRepo,
	}
}

func (euc *EventUseCase) CreatEvent(newEvent _entities.Event, idToken uint) (_entities.Event, error) {
	// UserId adalah id yang login
	newEvent.UserId = idToken

	// hosted by adalah nama user
	user, _, errGetUser := euc.userRepository.GetUser(int(idToken))
	if errGetUser != nil {
		return newEvent, errGetUser
	}
	newEvent.HostedBy = user.Name

	createEvent, err := euc.eventRepository.CreatEvent(newEvent)

	if newEvent.CatagoryId == 0 {
		return newEvent, errors.New("catagory_id can't be empty")
	}
	if newEvent.NameEvent == "" {
		return newEvent, errors.New("name_event can't be empty")
	}
	if newEvent.MaxParticipants == 0 {
		return newEvent, errors.New("max_participants can't be empty")
	}
	if newEvent.Location == "" {
		return newEvent, errors.New("location can't be empty")
	}
	if newEvent.DetailEvent == "" {
		return newEvent, errors.New("detail_event can't be empty")
	}
	if newEvent.UrlImage == "" {
		return newEvent, errors.New("url_image can't be empty")
	}
	if newEvent.Date.IsZero() {
		return newEvent, errors.New("date can't be empty")
	}

	return createEvent, err
}

func (euc *EventUseCase) GetAllEvent() ([]_entities.Event, error) {
	events, err := euc.eventRepository.GetAllEvent()
	return events, err
}

func (euc *EventUseCase) GetEventById(id int) (_entities.Event, int, error) {
	event, rows, err := euc.eventRepository.GetEventById(id)
	return event, rows, err
}

func (euc *EventUseCase) UpdateEvent(updateEvent _entities.Event, id int, idToken int) (_entities.Event, int, error) {
	eventFind, rows, err := euc.eventRepository.GetEventById(id)
	if err != nil {
		return eventFind, 0, err
	}
	if rows == 0 {
		return eventFind, 0, nil
	}
	if eventFind.UserId != uint(idToken) {
		return eventFind, 1, errors.New("unauthorized")
	}
	if updateEvent.CatagoryId != 0 {
		eventFind.CatagoryId = updateEvent.CatagoryId
	}
	if updateEvent.NameEvent != "" {
		eventFind.NameEvent = updateEvent.NameEvent
	}
	if updateEvent.MaxParticipants != 0 {
		eventFind.MaxParticipants = updateEvent.MaxParticipants
	}
	if updateEvent.Date.IsZero() {
		eventFind.Date = updateEvent.Date
	}
	if updateEvent.Location != "" {
		eventFind.Location = updateEvent.Location
	}
	if updateEvent.DetailEvent != "" {
		eventFind.DetailEvent = updateEvent.DetailEvent
	}
	if updateEvent.UrlImage != "" {
		eventFind.UrlImage = updateEvent.UrlImage
	}

	event, rows, err := euc.eventRepository.UpdateEvent(eventFind)
	return event, rows, err
}

func (euc *EventUseCase) DeleteEvent(id int) (int, error) {
	rows, err := euc.eventRepository.DeleteEvent(id)
	return rows, err
}
