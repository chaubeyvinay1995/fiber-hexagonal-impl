package zerologImpl

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

// Info This should be called for tracking the events happened in the application.
func Info(ctx *fiber.Ctx, message string) {
	log.Info().Msg(message)
}
