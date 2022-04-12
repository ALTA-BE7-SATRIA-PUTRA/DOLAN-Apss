package attendees

import (
	"fmt"
	_entities "group-project/dolan-planner/entities"

	"gorm.io/gorm"
)

type AttendeesRepository struct {
	database *gorm.DB
}

func NewAttendeesRepository(db *gorm.DB) *AttendeesRepository {
	return &AttendeesRepository{
		database: db,
	}
}

func (ur *AttendeesRepository) PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, int, error) {
	var attendees _entities.Attendees
	var event _entities.Event

	txEvent := ur.database.Where("id = ?", idEvent).Find(&event)

	if txEvent.Error != nil {
		return _entities.Attendees{}, 1, fmt.Errorf("fail to read event")
	}

	if event.TotalParticipants == event.MaxParticipants {
		return _entities.Attendees{}, 2, fmt.Errorf("quota full")
	}

	var attendeesdb []_entities.Attendees
	txAtt := ur.database.Where("event_id = ?", idEvent).Where("user_id = ?", idToken).Find(&attendeesdb)

	if txAtt.RowsAffected > 0 {
		return _entities.Attendees{}, 3, fmt.Errorf("you have join")
	}

	if txAtt.Error != nil {
		return _entities.Attendees{}, 4, fmt.Errorf("fail to read attendees")
	}

	attendees.UserId = idToken
	attendees.EventId = idEvent
	tx := ur.database.Save(&attendees)
	ur.database.Exec("UPDATE events SET total_participants = ? WHERE id = ?", gorm.Expr("total_participants + ?", 1), idEvent)

	if tx.Error != nil {
		return attendees, 5, tx.Error
	}
	return attendees, 0, nil
}

func (ur *AttendeesRepository) GetAttendees(idEvent uint) ([]_entities.Attendees, error) {
	var attendees []_entities.Attendees
	tx := ur.database.Where("event_id = ?", idEvent).Find(&attendees)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return attendees, nil
}
