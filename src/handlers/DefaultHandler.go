package handlers

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/utils"
)

type DefaultHandler struct {
	Logger         *core.Logger
	ResponseWriter *utils.ResponseWriter
}

func NewDefaultHandler(
	logger *core.Logger,
	responseWriter *utils.ResponseWriter,
) *DefaultHandler {
	return &DefaultHandler{
		Logger:         logger,
		ResponseWriter: responseWriter,
	}
}

func (defaultHandler *DefaultHandler) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defaultHandler.ResponseWriter.WriteCode(w, http.StatusOK)
	})
}
