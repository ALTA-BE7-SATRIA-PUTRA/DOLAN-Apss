package event

import (
	_entities "group-project/dolan-planner/entities"
)

type EventRepositoryInterface interface {
	CreatEvent(newEvent _entities.Event) (_entities.Event, error)
	GetAllEvent() ([]_entities.Event, error)
}
