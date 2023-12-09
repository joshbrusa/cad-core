package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/cad-http/src/handlers"
	"github.com/joshbrusa/cad-http/src/middlewares"
)

type DefaultRouter struct {
	Router *mux.Router
}

func NewDefaultRouter(router *mux.Router) *DefaultRouter {
	return &DefaultRouter{
		Router: router,
	}
}

func (defaultRouter *DefaultRouter) UseMiddlewares(
	loggerMiddleware *middlewares.LoggerMiddleware,
	panicMiddleware *middlewares.PanicMiddleware,
) {
	defaultRouter.Router.Use(loggerMiddleware.Handle)
	defaultRouter.Router.Use(panicMiddleware.Handle)
}

func (defaultRouter *DefaultRouter) HandleRoutes(defaultHandler *handlers.DefaultHandler) {
	defaultRouter.Router.Handle("/", defaultHandler.Handle())
}
