package main

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/router"
)

func main() {
	mux := http.NewServeMux()
	logger := core.NewLogger()

	env, envErr := core.NewEnv(logger)
	if envErr != nil {
		logger.Error(envErr.Error())
		return
	}

	postgres, postgresErr := core.NewPostgres(env, logger)
	if postgresErr != nil {
		logger.Error(postgresErr.Error())
		return
	}

	router := router.NewRouter(logger, mux, handler)
	router.Route()

	logger.Info("server listening", "port", env.Port)
	listenErr := http.ListenAndServe(":"+env.Port, mux)
	if listenErr != nil {
		logger.Error(listenErr.Error())
	}
}
