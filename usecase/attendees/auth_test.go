package attendees

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPostAttendees(t *testing.T) {
// 	t.Run("TestPostAttendeesSuccess", func(t *testing.T) {
// 		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepository{})
// 		data, idErr, err := attendeesUseCase.PostAttendees(idToken, idEvent)
// 		assert.Equal(t, "Token", data)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("TestPostAttendeesError", func(t *testing.T) {
// 		attendeesUseCase := NewAttendeesUseCase(mockAttendeesRepositoryError{})
// 		data, err := attendeesUseCase.PostAttendees(1, 2)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "", data)
// 	})
// }

// type mockAttendeesRepository struct{}

// func (m mockAttendeesRepository) PostAttendees(idEvent uint, idToken uint) (string, error) {
// 	return "Token", nil
// }

// func (m mockAttendeesRepository) GetAttendees(idEvent uint, idToken uint) (string, error) {
// 	return "Token", nil
// }

// func (m mockAttendeesRepository) DeleteAttendees(idEvent uint, idToken uint) (string, error) {
// 	return "Token", nil
// }

// type mockAttendeesRepositoryError struct{}

// func (m mockAttendeesRepositoryError) PostAttendees(email string, password string) (string, error) {
// 	return "", fmt.Errorf("Unautorizad")
// }

// func (m mockAttendeesRepositoryError) GetAttendees(email string, password string) (string, error) {
// 	return "", fmt.Errorf("Unautorizad")
// }

// func (m mockAttendeesRepositoryError) DeleteAttendees(email string, password string) (string, error) {
// 	return "", fmt.Errorf("Unautorizad")
// }
