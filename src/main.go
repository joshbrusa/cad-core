package main

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/routers"
)

func main() {
	// get logger
	logger := core.NewLogger()

	// get env
	env, envErr := core.NewEnv(logger)
	if envErr != nil {
		logger.Error(envErr.Error())
		return
	}

	// get postgres
	postgres, postgresErr := core.NewPostgres(env, logger)
	if postgresErr != nil {
		logger.Error(postgresErr.Error())
		return
	}

	// get router
	router := routers.NewRouter(env, logger, postgres)

	// listen
	logger.Info("server listening", "port", env.Port)
	listenErr := http.ListenAndServe(":"+env.Port, router)
	if listenErr != nil {
		logger.Error(listenErr.Error())
	}
}
