package attendees

import (
	_entities "group-project/dolan-planner/entities"
)

type AttendeesUseCaseInterface interface {
	PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, error)
	GetAttendees(idEvent uint) ([]_entities.Attendees, int, error)
	DeleteAttendees(idToken uint, idEvent uint) (uint, error)
	GetAttendeesUser(idToken uint) ([]_entities.Attendees, int, error)
}
