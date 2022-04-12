package event

import (
	_entities "group-project/dolan-planner/entities"

	"gorm.io/gorm"
)

type EventRepository struct {
	database *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{
		database: db,
	}
}

func (er *EventRepository) CreatEvent(newEvent _entities.Event) (_entities.Event, error) {
	tx := er.database.Save(&newEvent)
	if tx.Error != nil {
		return newEvent, tx.Error
	}
	return newEvent, nil
}

func (er *EventRepository) GetAllEvent() ([]_entities.Event, error) {
	var events []_entities.Event
	tx := er.database.Find(&events)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return events, nil
}

func (er *EventRepository) GetEventById(id int) (_entities.Event, int, error) {
	var event _entities.Event
	tx := er.database.Preload("Attendees.User").Preload("Comment.User").Find(&event, id)
	if tx.Error != nil {
		return event, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return event, 0, nil
	}
	return event, int(tx.RowsAffected), nil
}

func (er *EventRepository) UpdateEvent(updateEvent _entities.Event) (_entities.Event, int, error) {
	tx := er.database.Save(&updateEvent)
	if tx.Error != nil {
		return updateEvent, 0, tx.Error
	}
	return updateEvent, int(tx.RowsAffected), nil
}

func (er *EventRepository) DeleteEvent(id int) (int, error) {
	var event _entities.Event
	tx := er.database.Delete(&event, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
