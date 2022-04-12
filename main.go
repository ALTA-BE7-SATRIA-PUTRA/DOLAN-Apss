package main

import (
	"fmt"
	"group-project/dolan-planner/configs"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authHandler "group-project/dolan-planner/delivery/handler/auth"
	_authRepository "group-project/dolan-planner/repository/auth"
	_authUseCase "group-project/dolan-planner/usecase/auth"

	_userHandler "group-project/dolan-planner/delivery/handler/user"
	_userRepository "group-project/dolan-planner/repository/user"
	_userUseCase "group-project/dolan-planner/usecase/user"

	_eventHandler "group-project/dolan-planner/delivery/handler/event"
	_eventRepository "group-project/dolan-planner/repository/event"
	_eventUseCase "group-project/dolan-planner/usecase/event"

	_routes "group-project/dolan-planner/delivery/routes"
	_utils "group-project/dolan-planner/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	eventRepo := _eventRepository.NewEventRepository(db)
	eventUseCase := _eventUseCase.NewEventUseCase(eventRepo, userRepo)
	eventHandler := _eventHandler.NewEventHandler(eventUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterUserPath(e, userHandler)
	_routes.RegisterEventPath(e, eventHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
