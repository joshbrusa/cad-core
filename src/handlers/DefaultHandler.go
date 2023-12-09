package handlers

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/utils"
)

type DefaultHandler struct {
	Logger *core.Logger
}

func NewDefaultHandler(logger *core.Logger) *DefaultHandler {
	return &DefaultHandler{
		Logger: logger,
	}
}

func (defaultHandler *DefaultHandler) Handle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := "hello world"
		response := utils.NewResponse(w, http.StatusOK, msg)
		response.Send()
	})
}
