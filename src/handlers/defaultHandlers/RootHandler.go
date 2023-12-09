package defaultHandlers

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
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
		w.WriteHeader(http.StatusOK)
	})
}
