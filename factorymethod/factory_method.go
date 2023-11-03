package factorymethod

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("File:", message)
}

type LoggerFactory interface {
	CreateLogger() Logger
}

type ConsoleLoggerFactory struct{}

func (clf *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}

type FileLoggerFactory struct{}

func (flf *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

type EnvLoggerFactory struct{}

func (elf *EnvLoggerFactory) CreateLogger() Logger {
	loggerType := os.Getenv("LOGGER_TYPE")

	var logger Logger
	switch loggerType {
	case "console":
		logger = &ConsoleLogger{}
	case "file":
		logger = &FileLogger{}
	default:
		logger = &ConsoleLogger{}
	}

	return logger
}
