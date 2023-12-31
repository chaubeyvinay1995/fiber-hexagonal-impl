package api

import (
	"github.com/gofiber/fiber/v2"
	"hexagonal-fiber-impl/common/logger"
)

func SuccessResponse(ctx *fiber.Ctx, code int, data interface{}, message string) error {
	// Set the response status code
	ctx.Status(code)
	// Set the response content type to JSON
	ctx.Set("Content-Type", "application/json")
	// Create a response map
	response := map[string]interface{}{
		"status":  code,
		"message": message,
		"data":    data,
	}
	// Send the response as JSON
	return ctx.JSON(response)
}

func ErrorResponse(ctx *fiber.Ctx, code int, error interface{}, message string) error {
	// Set the response status code
	ctx.Status(code)
	// Set the response content type to JSON
	ctx.Set("Content-Type", "application/json")
	// Create a response map
	response := map[string]interface{}{
		"status":  code,
		"message": message,
		"error":   error,
	}
	logger.Error(ctx, error)
	// Send the response as JSON
	return ctx.JSON(response)
}
