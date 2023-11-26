package logger

import (
	"log/slog"
	"os"
)

func New() *Slog {
	return &Slog{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}
