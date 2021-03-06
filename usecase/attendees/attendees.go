package attendees

import (
	_entities "group-project/dolan-planner/entities"
	_attendeesRepository "group-project/dolan-planner/repository/attendees"
)

type AttendeesUseCase struct {
	attendeesRepository _attendeesRepository.AttendeesRepositoryInterface
}

func NewAttendeesUseCase(attendeesRepo _attendeesRepository.AttendeesRepositoryInterface) AttendeesUseCaseInterface {
	return &AttendeesUseCase{
		attendeesRepository: attendeesRepo,
	}
}

func (auc *AttendeesUseCase) PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, error) {
	attendeess, err := auc.attendeesRepository.PostAttendees(idEvent, idToken)
	return attendeess, err
}
func (auc *AttendeesUseCase) GetAttendees(idEvent uint) ([]_entities.Attendees, int, error) {
	attendeess, rows, err := auc.attendeesRepository.GetAttendees(idEvent)
	return attendeess, rows, err
}

func (auc *AttendeesUseCase) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	rows, err := auc.attendeesRepository.DeleteAttendees(idToken, idEvent)
	return rows, err
}

func (auc *AttendeesUseCase) GetAttendeesUser(idToken uint) ([]_entities.Attendees, int, error) {
	attendees, rows, err := auc.attendeesRepository.GetAttendeesUser(idToken)
	return attendees, rows, err
}
