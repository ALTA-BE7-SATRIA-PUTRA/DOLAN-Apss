package comment

import (
	_entities "group-project/dolan-planner/entities"
)

type CommentUseCaseInterface interface {
	PostComment(comment _entities.Comment, idEvent uint, idToken uint) (_entities.Comment, int, error)
	GetComment(idEvent uint) ([]_entities.Comment, error)
}
