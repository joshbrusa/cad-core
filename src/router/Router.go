package router

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/handler"
	"github.com/joshbrusa/cad-http/src/logger"
)

type Router struct {
	Logger  *logger.Logger
	Mux     *http.ServeMux
	Handler *handler.Handler
}

func NewRouter(
	logger *logger.Logger,
	mux *http.ServeMux,
	handler *handler.Handler,
) *Router {
	return &Router{
		Logger:  logger,
		Mux:     mux,
		Handler: handler,
	}
}
