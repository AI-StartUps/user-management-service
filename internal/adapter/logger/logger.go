package logger

import (
	"log"
	"os"
)

type DefaultLogger struct {
	consoleLogger *log.Logger
	fileLogger    *log.Logger
}

func NewDefaultLogger(logFilePath string) (*DefaultLogger, error) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &DefaultLogger{
		consoleLogger: log.New(os.Stdout, "", log.LstdFlags),
		fileLogger:    log.New(file, "", log.LstdFlags),
	}, nil
}

func (l *DefaultLogger) Info(message string) {
	l.consoleLogger.Println("[INFO]", message)
	l.fileLogger.Println("[INFO]", message)
}

func (l *DefaultLogger) Warning(message string) {
	l.consoleLogger.Println("[WARNING]", message)
	l.fileLogger.Println("[WARNING]", message)
}

func (l *DefaultLogger) Error(message string) {
	l.consoleLogger.Println("[ERROR]", message)
	l.fileLogger.Println("[ERROR]", message)
}
