package main

import (
	"fmt"
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
)

func main() {
	logger := core.NewLogger()
	env, envErr := core.NewEnv(logger)

	if envErr != nil {
		logger.Error(envErr.Error())
		return
	}

	mux := http.NewServeMux()
	// postgres := core.NewPostgres(logger, env)
	// handler := handler.NewHandler(logger)
	// router := router.NewRouter(logger, mux, handler)

	// router.RouteRoute()

	fmt.Println("server listening on port:", env.Port)
	listenErr := http.ListenAndServe(":"+env.Port, mux)

	if listenErr != nil {
		logger.Error(listenErr.Error())
	}
}
