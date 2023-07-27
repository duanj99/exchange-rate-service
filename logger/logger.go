package logger

import (
	"log"
	"os"
)

type ServiceLogger struct {
	logger *log.Logger
}

func NewLogger() *ServiceLogger {
	return &ServiceLogger{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

func (sl *ServiceLogger) Info(info string) {
	// Initialize a new logger which writes messages to the standard out stream,
	// prefixed with the current date and time.
	sl.logger.Print(info)
}

func (sl *ServiceLogger) Fatal(info string) {
	// Initialize a new logger which writes messages to the standard out stream,
	// prefixed with the current date and time.
	sl.logger.Fatal(info)
}
