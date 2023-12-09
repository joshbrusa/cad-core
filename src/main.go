package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/handlers"
	"github.com/joshbrusa/cad-http/src/middlewares"
	"github.com/joshbrusa/cad-http/src/routers"
)

func main() {
	// logger
	logger := core.NewLogger()

	// env
	env, envErr := core.NewEnv(logger)
	if envErr != nil {
		logger.Error(envErr.Error())
		return
	}

	// postgres
	_, postgresErr := core.NewPostgres(env, logger)
	if postgresErr != nil {
		logger.Error(postgresErr.Error())
		return
	}

	// middlewares
	loggerMiddleware := middlewares.NewLoggerMiddleware(logger)
	panicMiddleware := middlewares.NewPanicMiddleware(logger)

	// handlers
	defaultHandler := handlers.NewDefaultHandler(logger)

	// routers
	router := mux.NewRouter()

	defaultRouter := routers.NewDefaultRouter(router)
	defaultRouter.UseMiddlewares(loggerMiddleware, panicMiddleware)
	defaultRouter.HandleRoutes(defaultHandler)

	userRouter := routers.NewUserRouter(router)
	userRouter.UseMiddlewares(loggerMiddleware, panicMiddleware)

	// listen
	logger.Info("server listening", "port", env.Port)
	listenErr := http.ListenAndServe(":"+env.Port, router)
	if listenErr != nil {
		logger.Error(listenErr.Error())
	}
}
