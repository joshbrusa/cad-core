package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/handler"
	"github.com/joshbrusa/cad-http/src/router"
)

func main() {
	logger := core.NewLogger()
	env := core.NewEnv(logger)
	postgres := core.NewPostgres(logger, env)
	mux := http.NewServeMux()
	handler := handler.NewHandler(logger)
	router := router.NewRouter(logger, mux, handler)

	router.RouteRoute()

	fmt.Println("server listening on port:", env.Port)
	err := http.ListenAndServe(":"+env.Port, mux)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
