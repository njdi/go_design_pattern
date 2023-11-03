package factorymethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	loggerType := "console"

	var loggerFactory LoggerFactory
	switch loggerType {
	case "console":
		loggerFactory = &ConsoleLoggerFactory{}
	case "file":
		loggerFactory = &FileLoggerFactory{}
	default:
		panic("Invalid logger type")
	}

	logger := loggerFactory.CreateLogger()
	logger.Log("Hello World!")
}

func TestLogger(t *testing.T) {
	loggerType := "console"

	var logger Logger
	switch loggerType {
	case "console":
		logger = &ConsoleLogger{}
	case "file":
		logger = &FileLogger{}
	default:
		panic("Invalid logger type")
	}

	logger.Log("Hello World!")
}

func TestEnvLogger(t *testing.T) {
	loggerFactory := &EnvLoggerFactory{}

	logger := loggerFactory.CreateLogger()
	logger.Log("Hello World!")
}
