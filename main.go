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

	_attendeesHandler "group-project/dolan-planner/delivery/handler/attendees"
	_attendeesRepository "group-project/dolan-planner/repository/attendees"
	_attendeesUseCase "group-project/dolan-planner/usecase/attendees"

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

	attendeesRepo := _attendeesRepository.NewAttendeesRepository(db)
	attendeesUseCase := _attendeesUseCase.NewAttendeesUseCase(attendeesRepo)
	attendeesHandler := _attendeesHandler.NewAttendeesHandler(attendeesUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterUserPath(e, userHandler)
	_routes.RegisterAttendeesPath(e, attendeesHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
