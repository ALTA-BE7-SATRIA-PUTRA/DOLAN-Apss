package attendees

import (
	"fmt"
	"testing"

	_entities "group-project/dolan-planner/entities"

	"github.com/stretchr/testify/assert"
)

func TestPostAttendees(t *testing.T) {
	t.Run("TestPostAttendeesSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		attendees, err := attendeesUseCase.PostAttendees(1, 1)
		assert.Equal(t, uint(1), attendees.EventId)
		assert.Nil(t, err)
	})

	t.Run("TestPostAttendeesError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		_, err := attendeesUseCase.PostAttendees(1, 2)
		assert.NotNil(t, err)
	})
}
func TestGetAttendees(t *testing.T) {
	t.Run("TesGetAttendeesSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		attendees, rows, err := attendeesUseCase.GetAttendees(1)
		assert.Equal(t, uint(1), attendees[0].EventId)
		assert.Equal(t, 1, rows)
		assert.Nil(t, err)
	})

	t.Run("TestGetAttendeesError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		_, _, err := attendeesUseCase.GetAttendees(1)
		assert.NotNil(t, err)
	})
}
func TestDeleteAttendees(t *testing.T) {
	t.Run("TestDeleteAttendeesSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		rows, err := attendeesUseCase.DeleteAttendees(1, 1)
		assert.Equal(t, uint(1), rows)
		assert.Nil(t, err)
	})

	t.Run("TestDeleteAttendeesError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		_, err := attendeesUseCase.DeleteAttendees(1, 1)
		assert.NotNil(t, err)
	})
}
func TestGetAttendeesUser(t *testing.T) {
	t.Run("TestGetAttendeesUserSuccess", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
		attendees, rows, err := attendeesUseCase.GetAttendeesUser(1)
		assert.Equal(t, uint(1), attendees[0].EventId)
		assert.Equal(t, 1, rows)
		assert.Nil(t, err)
	})

	t.Run("TestGetAttendeesUserError", func(t *testing.T) {
		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
		_, _, err := attendeesUseCase.GetAttendeesUser(1)
		assert.NotNil(t, err)
	})
}

type mockAttendeesRepository struct{}

func (m mockAttendeesRepository) PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, error) {
	return _entities.Attendees{
		EventId: 1, UserId: 1,
	}, nil
}

func (m mockAttendeesRepository) GetAttendees(idEvent uint) ([]_entities.Attendees, int, error) {
	return []_entities.Attendees{
		{EventId: 1, UserId: 1},
	}, 1, nil
}

func (m mockAttendeesRepository) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	return 1, nil
}

func (m mockAttendeesRepository) GetAttendeesUser(idToken uint) ([]_entities.Attendees, int, error) {
	return []_entities.Attendees{
		{EventId: 1, UserId: 1},
	}, 1, nil
}

type mockAttendeesRepositoryError struct{}

func (m mockAttendeesRepositoryError) PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, error) {
	return _entities.Attendees{}, fmt.Errorf("failed to post attendees")
}

func (m mockAttendeesRepositoryError) GetAttendees(idEvent uint) ([]_entities.Attendees, int, error) {
	return nil, 0, fmt.Errorf("failed to get attendes")
}

func (m mockAttendeesRepositoryError) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	return 0, fmt.Errorf("failed to delete attendes")
}

func (m mockAttendeesRepositoryError) GetAttendeesUser(idToken uint) ([]_entities.Attendees, int, error) {
	return nil, 1, fmt.Errorf("failed to get attendes by user")
}
