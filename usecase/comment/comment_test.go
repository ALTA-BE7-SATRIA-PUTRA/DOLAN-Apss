package comment

import (
	"fmt"
	"testing"

	_entities "group-project/dolan-planner/entities"

	"github.com/stretchr/testify/assert"
)

func TestPostComment(t *testing.T) {
	t.Run("TestPostCommentSuccess", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepository{})
		comment, idErr, err := commentUseCase.PostComment(_entities.Comment{}, 1, 1)
		assert.Equal(t, "comment", comment.Comment)
		assert.Equal(t, 0, idErr)
		assert.Nil(t, err)
	})

	t.Run("TestPostCommentError", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepositoryError{})
		_, idErr, err := commentUseCase.PostComment(_entities.Comment{Comment: "comment"}, 1, 1)
		assert.Equal(t, 1, idErr)
		assert.NotNil(t, err)
	})
}
func TestGetComment(t *testing.T) {
	t.Run("TesGetCommentSuccess", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepository{})
		comment, err := commentUseCase.GetComment(1)
		assert.Equal(t, uint(1), comment[0].EventId)
		assert.Nil(t, err)
	})

	t.Run("TestGetCommentError", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepositoryError{})
		_, err := commentUseCase.GetComment(1)
		assert.NotNil(t, err)
	})
}

type mockCommentRepository struct{}

func (m mockCommentRepository) PostComment(comment _entities.Comment, idEvent uint, idToken uint) (_entities.Comment, int, error) {
	return _entities.Comment{
		EventId: 1, UserId: 1, Comment: "comment",
	}, 0, nil
}

func (m mockCommentRepository) GetComment(idEvent uint) ([]_entities.Comment, error) {
	return []_entities.Comment{
		{EventId: 1, UserId: 1},
	}, nil
}

type mockCommentRepositoryError struct{}

func (m mockCommentRepositoryError) PostComment(comment _entities.Comment, idEvent uint, idToken uint) (_entities.Comment, int, error) {
	return _entities.Comment{}, 1, fmt.Errorf("failed to post comment in to event with id event")
}

func (m mockCommentRepositoryError) GetComment(idEvent uint) ([]_entities.Comment, error) {
	return nil, fmt.Errorf("failed to get comment by id event")
}
