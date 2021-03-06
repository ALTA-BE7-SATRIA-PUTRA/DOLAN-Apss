package routes

import (
	_attendeesHandler "group-project/dolan-planner/delivery/handler/attendees"
	_authHandler "group-project/dolan-planner/delivery/handler/auth"
	_catagoryHandler "group-project/dolan-planner/delivery/handler/catagory"
	_commentHandler "group-project/dolan-planner/delivery/handler/comment"
	_eventHandler "group-project/dolan-planner/delivery/handler/event"
	_userHandler "group-project/dolan-planner/delivery/handler/user"
	_middlewares "group-project/dolan-planner/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/login", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}
func RegisterAttendeesPath(e *echo.Echo, ah *_attendeesHandler.AttendeesHandler) {
	e.POST("/attendees/:id", ah.PostAttendeesHandler(), _middlewares.JWTMiddleware())
	e.GET("/attendees/:id", ah.GetAttendeesHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/attendees/:id", ah.DeleteAttendeesHandler(), _middlewares.JWTMiddleware())
	e.GET("/attendees/user", ah.GetAttendeesUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterEventPath(e *echo.Echo, eh *_eventHandler.EventHandler) {
	e.POST("/events", eh.CreateEventHandler(), _middlewares.JWTMiddleware())
	e.GET("/events", eh.GetAllEventHandler())
	e.GET("/events/:id", eh.GetEventByIdHandler())
	e.GET("/events/user", eh.GetEventByUserIdHandler(), _middlewares.JWTMiddleware())
	e.PUT("/events/:id", eh.UpdateEventHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/events/:id", eh.DeleteEventHandler(), _middlewares.JWTMiddleware())
}

func RegisterCatagoryPath(e *echo.Echo, uh *_catagoryHandler.CatagoryHandler) {
	e.GET("/catagories", uh.GetAllCatagoryHandler())
	e.POST("/catagories", uh.CreateCatagoryHandler())
}

func RegisterCommentPath(e *echo.Echo, ch *_commentHandler.CommentHandler) {
	e.POST("/comment/:id", ch.PostCommentHandler(), _middlewares.JWTMiddleware())
	e.GET("/comment/:id", ch.GetCommentHandler(), _middlewares.JWTMiddleware())
}
