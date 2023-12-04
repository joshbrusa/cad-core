package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
