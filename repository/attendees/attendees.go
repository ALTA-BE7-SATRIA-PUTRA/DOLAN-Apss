package attendees

import (
	"errors"
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

func (ur *AttendeesRepository) PostAttendees(idEvent uint, idToken uint) (_entities.Attendees, error) {
	var attendees _entities.Attendees
	var event _entities.Event

	txEvent := ur.database.Where("id = ?", idEvent).Find(&event)
	if txEvent.RowsAffected == 0 {
		return _entities.Attendees{}, fmt.Errorf("not found")
	}

	if txEvent.Error != nil {
		return _entities.Attendees{}, fmt.Errorf("fail to read event")
	}

	if event.TotalParticipants == event.MaxParticipants {
		return _entities.Attendees{}, fmt.Errorf("quota full")
	}

	var attendeesdb []_entities.Attendees
	txAtt := ur.database.Where("event_id = ?", idEvent).Where("user_id = ?", idToken).Find(&attendeesdb)
	if txAtt.RowsAffected == 1 {
		return _entities.Attendees{}, fmt.Errorf("you have join")
	}

	if txAtt.Error != nil {
		return _entities.Attendees{}, fmt.Errorf("fail to read attendees")
	}

	attendees.UserId = idToken
	attendees.EventId = idEvent
	tx := ur.database.Save(&attendees)
	ur.database.Exec("UPDATE events SET total_participants = ? WHERE id = ?", gorm.Expr("total_participants + ?", 1), idEvent)
	if tx.Error != nil {
		return attendees, tx.Error
	}

	return attendees, nil
}

func (ur *AttendeesRepository) GetAttendees(idEvent uint) ([]_entities.Attendees, int, error) {
	var attendees []_entities.Attendees
	tx := ur.database.Preload("User").Where("event_id = ?", idEvent).Find(&attendees)

	if tx.Error != nil {
		return nil, 1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("not found")
	}
	return attendees, 1, nil
}

func (ar *AttendeesRepository) DeleteAttendees(idToken uint, idEvent uint) (uint, error) {
	var attendees _entities.Attendees
	tx := ar.database.Where("event_id = ?", idEvent).Where("user_id = ?", idToken).Unscoped().Delete(&attendees)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}

	// mengurangi total participants karena ada user yang left
	ar.database.Exec("UPDATE events SET total_participants = ? WHERE id = ?", gorm.Expr("total_participants - ?", 1), idEvent)

	return uint(tx.RowsAffected), nil
}

func (ur *AttendeesRepository) GetAttendeesUser(idToken uint) ([]_entities.Attendees, int, error) {
	var attendees []_entities.Attendees
	tx := ur.database.Preload("Event").Where("user_id = ?", idToken).Find(&attendees)

	if tx.Error != nil {
		return attendees, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return attendees, 0, nil
	}
	return attendees, int(tx.RowsAffected), nil
}
