package api

import (
	"hexagonal-fiber-impl/core/domain"
	"hexagonal-fiber-impl/core/ports"

	fiber "github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.IUserService
}

var _ ports.IUserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.IUserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var payload *domain.UserLoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return ErrorResponse(c, 400, err.Error(), "Error while parsing request data.")
	}
	errors := payload.Validate()
	if errors != nil {
		return ErrorResponse(c, 400, errors, "Validation error")
	}
	user, err := h.userService.Login(payload.Email, payload.Password)

	if err != nil {
		return ErrorResponse(c, 400, err.Error(), "User not found")
	}

	return SuccessResponse(c, 200, user, "Logged in successfully")
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	var payload *domain.UserRegisterRequest
	if err := c.BodyParser(&payload); err != nil {
		return ErrorResponse(c, 400, err.Error(), "Error while parsing request data.")
	}
	errors := payload.Validate()
	if errors != nil {
		return ErrorResponse(c, 400, errors, "Validation error")
	}
	user, err := h.userService.Register(payload.Email, payload.Password, payload.ConfirmPassword)
	if err != nil {
		return ErrorResponse(c, 400, err.Error(), "Operation failed")
	}
	return SuccessResponse(c, 200, user, "User register in successfully")
}
