package routers

import (
	"github.com/gorilla/mux"
	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/middlewares"
)

type Router struct {
	Env      *core.Env
	Logger   *core.Logger
	Postgres *core.Postgres
	Router   *mux.Router
}

func NewRouter(
	env *core.Env,
	logger *core.Logger,
	postgres *core.Postgres,
) *Router {
	router := mux.NewRouter()
	loggerMiddleware := middlewares.NewLoggerMiddleware(logger)

	return &Router{
		Env:      env,
		Logger:   logger,
		Postgres: postgres,
		Router:   router,
	}
}

func (router *Router) UseMiddlewares() {
	loggerMiddleware := middlewares.LoggerMiddleware(router.Logger)
	router.Router.Use(middlewares.LoggerMiddleware(router.Logger))
	router.Router.Use(middlewares.PanicMiddleware(router.Logger))
}
