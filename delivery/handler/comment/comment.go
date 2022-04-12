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
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("you haven't written a message yet"))
		}

		_, idErr, _ := uh.commentUseCase.PostComment(newComment, uint(idEvent), uint(idToken))
		if idErr == 1 {
			c.JSON(http.StatusInternalServerError, helper.ResponseFailed("need join to event"))
		}

		if idErr == 2 {
			c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to post comment"))
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
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("fail to get comment"))
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("succses to get comment", comment))
	}
}
