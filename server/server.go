package server

import (
	"github.com/gofiber/fiber/v2"
	"hexagonal-fiber-impl/core/ports"
	"log"
)

type Server struct {
	//We will add every new Handler here
	userHandlers ports.IUserHandlers
	//middlewares ports.IMiddlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(uHandlers ports.IUserHandlers) *Server {
	return &Server{
		userHandlers: uHandlers,
		//paymentHandlers: pHandlers
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/login", s.userHandlers.Login)
	userRoutes.Post("/register", s.userHandlers.Register)
	log.Fatal(app.Listen(":5000"))
}
