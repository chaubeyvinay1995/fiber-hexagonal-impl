package main

import (
	"hexagonal-fiber-impl/api"
	"hexagonal-fiber-impl/core/services"
	"hexagonal-fiber-impl/repositories"
	"hexagonal-fiber-impl/server"
)

func main() {
	mongoConn := "mongodb+srv://vinpythondev:YXFMI4DSmGAoD6Ep@cluster1.t6w2vqw.mongodb.net/?retryWrites=true&w=majority"
	//repositories
	userRepository := repositories.NewUserRepository(mongoConn)
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := api.NewUserHandlers(userService)
	//server
	server.NewServer(userHandlers).Initialize()
}
