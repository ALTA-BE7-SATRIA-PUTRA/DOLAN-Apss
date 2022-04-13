package event

import (
	_entities "group-project/dolan-planner/entities"
)

type EventUseCaseInterface interface {
	CreatEvent(newEvent _entities.Event, idToken uint) (_entities.Event, error)
	GetAllEvent(filters map[string]string) ([]_entities.Event, error)
	GetEventById(id int) (_entities.Event, int, error)
	UpdateEvent(updateEvent _entities.Event, id int, idToken int) (_entities.Event, int, error)
	DeleteEvent(id int) (int, error)
}
