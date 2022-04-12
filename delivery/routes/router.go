package routes

import (
	_authHandler "group-project/dolan-planner/delivery/handler/auth"
	_eventHandler "group-project/dolan-planner/delivery/handler/event"
	_userHandler "group-project/dolan-planner/delivery/handler/user"
	_middlewares "group-project/dolan-planner/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterEventPath(e *echo.Echo, eh *_eventHandler.EventHandler) {
	e.POST("/events", eh.CreateEventHandler(), _middlewares.JWTMiddleware())
	e.GET("/events", eh.GetAllEventHandler())
	e.GET("/events/:id", eh.GetEventByIdHandler())
	e.PUT("/events/:id", eh.UpdateEventHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/events/:id", eh.DeleteEventHandler(), _middlewares.JWTMiddleware())
}
