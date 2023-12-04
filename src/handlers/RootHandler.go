package handlers

import (
	"net/http"

	"github.com/joshbrusa/cad-core/src/logger"
)

type RootHandler struct {
	Logger *logger.Logger
}

func NewRootHandler(logger *logger.Logger) *RootHandler {
	return &RootHandler{
		Logger: logger,
	}
}

func (rootHandler *RootHandler) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rootHandler.Logger.Info("handling root", "path", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	})
}
