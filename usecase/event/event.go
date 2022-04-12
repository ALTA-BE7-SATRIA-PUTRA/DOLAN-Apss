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

	return createEvent, err
}

func (euc *EventUseCase) GetAllEvent() ([]_entities.Event, error) {
	events, err := euc.eventRepository.GetAllEvent()
	return events, err
}
