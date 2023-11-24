package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshbrusa/cad-core/internal/handlers"
)

func (server *Server) Start() {
	server.Mux.Handle("/", handlers.RootHandler())

	fmt.Println("server listening on port: ", server.Port)
	err := http.ListenAndServe(":"+server.Port, server.Mux)

	if err != nil {
		server.Logger.Error(err.Error())
		os.Exit(1)
	}
}
