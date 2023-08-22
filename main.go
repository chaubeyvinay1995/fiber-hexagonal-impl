package main

import (
	"hexagonal-fiber-impl/api"
	"hexagonal-fiber-impl/core/services"
	"hexagonal-fiber-impl/repositories"
	"hexagonal-fiber-impl/server"
)

func main() {
	mongoConn := "secret🤫"
	//repositories
	userRepository := repositories.NewUserRepository(mongoConn)
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := api.NewUserHandlers(userService)
	//server
	server.NewServer(userHandlers).Initialize()
}
