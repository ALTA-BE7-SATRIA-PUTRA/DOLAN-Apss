package attendees

import (
	_entities "group-project/dolan-planner/entities"
)

type AttendeesRepositoryInterface interface {
	PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, int, error)
	GetAttendees(idEvent uint) ([]_entities.Attendees, error)
	DeleteAttendees(idToken uint, idEvent uint) (uint, error)
}
