package logger

import (
	"log/slog"
	"os"
)

func New() *Logger {
	return &Logger{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
