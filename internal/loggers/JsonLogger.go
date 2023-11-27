package loggers

import (
	"log/slog"
	"os"
)

type JsonLogger struct {
	*slog.Logger
}

func NewJsonLogger() *JsonLogger {
	return &JsonLogger{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
