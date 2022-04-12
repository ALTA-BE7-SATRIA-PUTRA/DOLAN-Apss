package event

import (
	_entities "group-project/dolan-planner/entities"
)

type EventUseCaseInterface interface {
	CreatEvent(newEvent _entities.Event, idToken uint) (_entities.Event, error)
	GetAllEvent() ([]_entities.Event, error)
}
