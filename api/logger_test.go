package api

import (
	fiber "github.com/gofiber/fiber/v2"
	"hexagonal-fiber-impl/common/logger"
	"hexagonal-fiber-impl/common/zerologImpl"
	"testing"
)

func BenchmarkLogrusPackage(b *testing.B) {
	fiberContext := fiber.Ctx{}
	for i := 0; i < b.N; i++ {
		logger.Info(&fiberContext, "Hello")
	}
}

func BenchmarkZeroLogPackage(b *testing.B) {
	fiberContext := fiber.Ctx{}
	for i := 0; i < b.N; i++ {
		zerologImpl.Info(&fiberContext, "hello")
	}
}
