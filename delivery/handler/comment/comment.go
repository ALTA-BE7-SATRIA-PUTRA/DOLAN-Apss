package comment

import (
	"group-project/dolan-planner/delivery/helper"
	_middlewares "group-project/dolan-planner/delivery/middlewares"
	_entities "group-project/dolan-planner/entities"
	_commentUseCase "group-project/dolan-planner/usecase/comment"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUseCase _commentUseCase.CommentUseCaseInterface
}

func NewCommentHandler(commentUseCase _commentUseCase.CommentUseCaseInterface) *CommentHandler {
	return &CommentHandler{
		commentUseCase: commentUseCase,
	}
}

func (uh *CommentHandler) PostCommentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newComment _entities.Comment
		c.Bind(&newComment)

		comment, errFil := helper.FilterComment(newComment.Comment)
		if errFil != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail in filter comment"))
		}
		newComment.Comment = comment

		idStr := c.Param("id")
		idEvent, errorconv := strconv.Atoi(idStr)
		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be number id event")
		}

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		if newComment.Comment == "" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("you haven't written a message yet or bad format"))
		}

		_, idErr, _ := uh.commentUseCase.PostComment(newComment, uint(idEvent), uint(idToken))
		if idErr == 1 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("need join to event"))
		}

		if idErr == 2 {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to post comment"))
		}

		if idErr == 3 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("event not found"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("succses to post comment"))

	}
}

func (uh *CommentHandler) GetCommentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		idEvent, errorconv := strconv.Atoi(idStr)
		if errorconv != nil {
			return c.JSON(http.StatusBadRequest, "The expected param must be number id event")
		}

		comment, err := uh.commentUseCase.GetComment(uint(idEvent))

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail to get comment"))
		}

		responseComment := []map[string]interface{}{}
		for i := 0; i < len(comment); i++ {
			response := map[string]interface{}{
				"id":         comment[i].ID,
				"created_at": comment[i].CreatedAt,
				"event_id":   comment[i].EventId,
				"user_id":    comment[i].UserId,
				"comment":    comment[i].Comment,
				"user": map[string]interface{}{
					"name": comment[i].User.Name},
			}
			responseComment = append(responseComment, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to get comment", responseComment))
	}
}
