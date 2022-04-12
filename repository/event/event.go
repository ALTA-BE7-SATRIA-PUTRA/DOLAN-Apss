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
