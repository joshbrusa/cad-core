package main

import (
	"fmt"
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/handler"
	"github.com/joshbrusa/cad-http/src/router"
)

func main() {
	logger := core.NewLogger()
	env, envErr := core.NewEnv(logger)

	if envErr != nil {
		logger.Error(envErr.Error())
		return
	}

	mux := http.NewServeMux()
	postgres := core.NewPostgres(logger, env)
	handler := handler.NewHandler(logger)
	router := router.NewRouter(logger, mux, handler)

	router.RouteRoute()

	fmt.Println("server listening on port:", "8000")
	listenErr := http.ListenAndServe(":"+"8000", mux)

	if listenErr != nil {
		logger.Error(listenErr.Error())
	}
}
