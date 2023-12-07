package handler

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
)

type RootHandler struct {
	Env    *core.Env
	Logger *core.Logger
}

func NewRootHandler(env *core.Env, logger *core.Logger) *RootHandler {
	return &RootHandler{
		Env:    env,
		Logger: logger,
	}
}

func (rootHandler *RootHandler) HandleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
