package ports

import (
	fiber "github.com/gofiber/fiber/v2"
	"hexagonal-fiber-impl/core/domain"
)

type IUserService interface {
	Login(email string, password string) (domain.User, error)
	Register(email string, password string, passwordConfirmation string) (domain.User, error)
}

type IUserRepository interface {
	Login(email string, password string) (domain.User, error)
	Register(email string, password string) (domain.User, error)
}

type IUserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
