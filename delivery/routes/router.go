package routes

import (
	_attendeesHandler "group-project/dolan-planner/delivery/handler/attendees"
	_authHandler "group-project/dolan-planner/delivery/handler/auth"
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
}
