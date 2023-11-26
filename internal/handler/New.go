package handler

import "github.com/joshbrusa/cad-core/internal/logger"

func New(slog *logger.Slog) *Handler {
	return &Handler{
		Slog: slog,
	}
}
