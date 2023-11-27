package handlers

import (
	"net/http"

	"github.com/joshbrusa/cad-core/internal/loggers"
)

type RootHandler struct {
	JsonLogger *loggers.JsonLogger
}

func NewRootHandler(jsonLogger *loggers.JsonLogger) *RootHandler {
	return &RootHandler{
		JsonLogger: jsonLogger,
	}
}

func (rootHandler *RootHandler) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rootHandler.JsonLogger.Info("handling root", "path", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	})
}
