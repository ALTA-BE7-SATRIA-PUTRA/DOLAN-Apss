package comment

import (
	"errors"
	"fmt"
	_entities "group-project/dolan-planner/entities"

	"gorm.io/gorm"
)

type CommentRepository struct {
	database *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		database: db,
	}
}

func (ur *CommentRepository) PostComment(comment _entities.Comment, idEvent uint, idToken uint) (_entities.Comment, int, error) {
	var attendeesdb _entities.Attendees
	var event _entities.Event

	comment.UserId = idToken
	comment.EventId = idEvent

	txEvent := ur.database.Where("id = ?", idEvent).Find(&event)

	if txEvent.RowsAffected == 0 {
		return _entities.Comment{}, 3, fmt.Errorf("not found")
	}

	txAtt := ur.database.Where("event_id = ?", idEvent).Where("user_id = ?", idToken).Find(&attendeesdb)
	if txAtt.RowsAffected == 0 {
		return _entities.Comment{}, 1, fmt.Errorf("need join to event")
	}
	tx := ur.database.Save(&comment)

	if tx.Error != nil {
		return comment, 2, tx.Error
	}
	return comment, 0, nil
}

func (ur *CommentRepository) GetComment(idEvent uint) ([]_entities.Comment, error) {
	var comment []_entities.Comment
	
	tx := ur.database.Preload("User").Where("event_id = ?", idEvent).Find(&comment)

  if tx.RowsAffected == 0 {
      return nil, errors.New("failed to get comment: event not found")
    }

	if tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}
