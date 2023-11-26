package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshbrusa/cad-core/internal/handler"
)

func (server *Server) Start() {
	handler := handler.New(server.Slog)

	server.Mux.Handle("/", handler.Root())

	fmt.Println("server listening on port:", server.Port)
	err := http.ListenAndServe(":"+server.Port, server.Mux)

	if err != nil {
		server.Slog.Error(err.Error())
		os.Exit(1)
	}
}
