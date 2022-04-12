package comment

import (
	_entities "group-project/dolan-planner/entities"
	_commentRepository "group-project/dolan-planner/repository/comment"
)

type CommentUseCase struct {
	commentRepository _commentRepository.CommentRepositoryInterface
}

func NewCommentUseCase(commentRepo _commentRepository.CommentRepositoryInterface) CommentUseCaseInterface {
	return &CommentUseCase{
		commentRepository: commentRepo,
	}
}

func (auc *CommentUseCase) PostComment(comment _entities.Comment, idEvent uint, idToken uint) (_entities.Comment, int, error) {
	comments, idErr, err := auc.commentRepository.PostComment(comment, idEvent, idToken)

	return comments, idErr, err
}
func (auc *CommentUseCase) GetComment(idEvent uint) ([]_entities.Comment, error) {
	comments, err := auc.commentRepository.GetComment(idEvent)
	return comments, err
}
