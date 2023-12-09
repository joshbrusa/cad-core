package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/handlers"
	"github.com/joshbrusa/cad-http/src/middlewares"
	"github.com/joshbrusa/cad-http/src/routers"
	"github.com/joshbrusa/cad-http/src/utils"
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

	// response writer
	responseWriter := utils.NewResponseWriter(logger)

	// middlewares
	loggerMiddleware := middlewares.NewLoggerMiddleware(logger)
	panicMiddleware := middlewares.NewPanicMiddleware(logger, responseWriter)

	// handlers
	defaultHandler := handlers.NewDefaultHandler(logger, responseWriter)

	// router
	router := mux.NewRouter()

	// default router
	defaultRouter := routers.NewDefaultRouter(router)
	defaultRouter.UseMiddlewares(loggerMiddleware, panicMiddleware)
	defaultRouter.HandleRoutes(defaultHandler)

	// user router
	userRouter := routers.NewUserRouter(router)
	userRouter.UseMiddlewares(loggerMiddleware, panicMiddleware)

	// listen
	logger.Info("server listening", "port", env.Port)
	listenErr := http.ListenAndServe(":"+env.Port, router)
	if listenErr != nil {
		logger.Error(listenErr.Error())
	}
}
