package handler

import (
	"github.com/joshbrusa/cad-http/src/logger"
)

type Handler struct {
	Logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
