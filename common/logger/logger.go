package logger

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

var logger = logrus.New()

const repoPath = ""

func init() {
	logger.SetReportCaller(true)
	logger.Formatter = getFormatter(0)
}

var getFormatter = func(skip int) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			defer func() { recover() }()
			_, fn, line, _ := runtime.Caller(skip)
			repoPath := fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), repoPath)
			fileName := strings.Replace(fn, repoPath, "", -1)
			return "", fmt.Sprintf("%s:%d", fileName, line)
		},
	}
}

// Info This should be called for tracking the events happened in the application.
func Info(ctx *fiber.Ctx, args ...interface{}) {
	logger.Info(ctx, args)
}

// Error function will be called while getting error.
func Error(c *fiber.Ctx, args ...interface{}) {
	logger.Error(c, args)
}

func Debug(ctx *fiber.Ctx, args ...interface{}) {
	logger.Debug(ctx, args)
}

func WithAlert(c *fiber.Ctx, err interface{}, args ...interface{}) {
	if err, ok := err.(error); ok {
		logger.Error(err)
	}
	Error(c, err, args)
}
