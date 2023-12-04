package handlers

import (
	"net/http"

	"github.com/joshbrusa/cad-core/src/core"
)

type RootHandler struct {
	Logger *core.Logger
}

func NewRootHandler(logger *core.Logger) *RootHandler {
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
