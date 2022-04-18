package event

import (
	"fmt"
	_entities "group-project/dolan-planner/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	t.Run("TestCreateEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{}, mockUserRepository{})
		layoutFormat := "2006-01-02T15:04:05Z0700"
		dateFormat := "2022-05-02T10:10:00+0700"
		dateParse, _ := time.Parse(layoutFormat, dateFormat)

		event, err := eventUseCase.CreatEvent(_entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"}, 1)

		assert.Nil(t, err)
		assert.Equal(t, "running", event.NameEvent)
	})

	t.Run("TestCreateEventError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{}, mockUserRepositoryError{})
		layoutFormat := "2006-01-02T15:04:05Z0700"
		dateFormat := "2022-05-02T10:10:00+0700"
		dateParse, _ := time.Parse(layoutFormat, dateFormat)

		event, err := eventUseCase.CreatEvent(_entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"}, 1)

		assert.NotNil(t, err)
		assert.Equal(t, "running", event.NameEvent)
	})
}

func TestGetAllEvent(t *testing.T) {
	t.Run("TestGetAllEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{}, mockUserRepository{})
		var filters map[string]string
		data, err := eventUseCase.GetAllEvent(filters)
		assert.Nil(t, err)
		assert.NotNil(t, data)
	})

	t.Run("TestGetAllEventError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{}, mockUserRepositoryError{})
		var filters map[string]string
		data, err := eventUseCase.GetAllEvent(filters)
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetEventById(t *testing.T) {
	t.Run("TestGetEventByIdSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{}, mockUserRepository{})
		event, rows, err := eventUseCase.GetEventById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "running", event.NameEvent)
	})

	t.Run("TestGetEventByIdError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{}, mockUserRepositoryError{})
		event, rows, err := eventUseCase.GetEventById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, "", event.NameEvent)
	})
}

func TestGetEventByUserId(t *testing.T) {
	t.Run("TestGetEventByUserIdSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{}, mockUserRepository{})
		events, rows, err := eventUseCase.GetEventByUserId(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.NotNil(t, events)
	})

	t.Run("TestGetEventByUserIdError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{}, mockUserRepositoryError{})
		events, rows, err := eventUseCase.GetEventByUserId(1)
		assert.NotNil(t, err)
		assert.Nil(t, events)
		assert.Equal(t, 0, rows)
	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("TestUpdateEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{}, mockUserRepository{})
		layoutFormat := "2006-01-02T15:04:05Z0700"
		dateFormat := "2022-05-02T10:10:00+0700"
		dateParse, _ := time.Parse(layoutFormat, dateFormat)
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"}, 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
		assert.Equal(t, "running", event.NameEvent)
	})

	t.Run("TestUpdateEventError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{}, mockUserRepositoryError{})
		layoutFormat := "2006-01-02T15:04:05Z0700"
		dateFormat := "2022-05-02T10:10:00+0700"
		dateParse, _ := time.Parse(layoutFormat, dateFormat)
		event, rows, err := eventUseCase.UpdateEvent(_entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"}, 1, 1)

		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, "", event.HostedBy)
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("TestDeleteEventSuccess", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepository{}, mockUserRepository{})
		rows, err := eventUseCase.DeleteEvent(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestDeleteEventError", func(t *testing.T) {
		eventUseCase := NewEventUseCase(mockEventRepositoryError{}, mockUserRepositoryError{})
		rows, err := eventUseCase.DeleteEvent(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
	})
}

// === mock success ===
type mockEventRepository struct{}

func (m mockEventRepository) CreatEvent(newEvent _entities.Event) (_entities.Event, error) {
	layoutFormat := "2006-01-02T15:04:05Z0700"
	dateFormat := "2022-05-02T10:10:00+0700"
	dateParse, _ := time.Parse(layoutFormat, dateFormat)
	return _entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"},
		nil
}

func (m mockEventRepository) GetAllEvent(filters map[string]string) ([]_entities.Event, error) {
	layoutFormat := "2006-01-02T15:04:05Z0700"
	dateFormat := "2022-05-02T10:10:00+0700"
	dateParse, _ := time.Parse(layoutFormat, dateFormat)
	return []_entities.Event{
			{UserId: 1,
				CatagoryId:        1,
				NameEvent:         "running",
				HostedBy:          "usamah",
				MaxParticipants:   50,
				TotalParticipants: 0,
				Date:              dateParse,
				Location:          "Bogor",
				DetailEvent:       "running-running",
				UrlImage:          "event.jgp"}},
		nil
}

func (m mockEventRepository) GetEventById(id int) (_entities.Event, int, error) {
	layoutFormat := "2006-01-02T15:04:05Z0700"
	dateFormat := "2022-05-02T10:10:00+0700"
	dateParse, _ := time.Parse(layoutFormat, dateFormat)
	return _entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"},
		1, nil
}

func (m mockEventRepository) GetEventByUserId(idToken uint) ([]_entities.Event, int, error) {
	layoutFormat := "2006-01-02T15:04:05Z0700"
	dateFormat := "2022-05-02T10:10:00+0700"
	dateParse, _ := time.Parse(layoutFormat, dateFormat)
	return []_entities.Event{
			{UserId: 1,
				CatagoryId:        1,
				NameEvent:         "running",
				HostedBy:          "usamah",
				MaxParticipants:   50,
				TotalParticipants: 0,
				Date:              dateParse,
				Location:          "Bogor",
				DetailEvent:       "running-running",
				UrlImage:          "event.jgp"}},
		1, nil
}

func (m mockEventRepository) UpdateEvent(updateEvent _entities.Event) (_entities.Event, int, error) {
	layoutFormat := "2006-01-02T15:04:05Z0700"
	dateFormat := "2022-05-02T10:10:00+0700"
	dateParse, _ := time.Parse(layoutFormat, dateFormat)
	return _entities.Event{
			UserId:            1,
			CatagoryId:        1,
			NameEvent:         "running",
			HostedBy:          "usamah",
			MaxParticipants:   50,
			TotalParticipants: 0,
			Date:              dateParse,
			Location:          "Bogor",
			DetailEvent:       "running-running",
			UrlImage:          "event.jgp"},
		1, nil
}

func (m mockEventRepository) DeleteEvent(id int) (int, error) {
	return 1, nil
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) CreatUser(createUser _entities.User) (_entities.User, error) {
	user := _entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}
	return user, nil
}

func (m mockUserRepository) GetUser(idToken int) (_entities.User, int, error) {
	return _entities.User{
		Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah",
	}, 1, nil
}

func (m mockUserRepository) UpdateUser(userUpdate _entities.User) (_entities.User, int, error) {
	user := _entities.User{Name: "Usamah", City: "Bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}
	return user, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) (int, error) {
	return 1, nil
}

// === mock error ===
type mockEventRepositoryError struct{}

func (m mockEventRepositoryError) CreatEvent(newEvent _entities.Event) (_entities.Event, error) {
	return _entities.Event{}, fmt.Errorf("create event error")
}

func (m mockEventRepositoryError) GetAllEvent(filters map[string]string) ([]_entities.Event, error) {
	return nil, fmt.Errorf("Get All Event Error")
}

func (m mockEventRepositoryError) GetEventById(id int) (_entities.Event, int, error) {
	return _entities.Event{}, 0, fmt.Errorf("Get Event By Id Error")
}

func (m mockEventRepositoryError) GetEventByUserId(idToken uint) ([]_entities.Event, int, error) {
	return nil, 0, fmt.Errorf("Get Event By User Id Error")
}

func (m mockEventRepositoryError) UpdateEvent(updateEvent _entities.Event) (_entities.Event, int, error) {
	return _entities.Event{}, 0, fmt.Errorf("Update Event Error")
}

func (m mockEventRepositoryError) DeleteEvent(id int) (int, error) {
	return 0, fmt.Errorf("Delete Event Error")
}

// === mock error ===
type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) CreatUser(user _entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create user")
}

func (m mockUserRepositoryError) GetUser(idToken int) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error get user")
}

func (m mockUserRepositoryError) UpdateUser(userUpdate _entities.User) (_entities.User, int, error) {
	return _entities.User{Name: "usamah", City: "bogor", UrlImage: "usamah.com", Email: "usamah@gmail.com", Password: "usamah"}, 0, fmt.Errorf("error update user")
}

func (m mockUserRepositoryError) DeleteUser(id int) (int, error) {
	return 0, fmt.Errorf("error delete user")
}
