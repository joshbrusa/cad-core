package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshbrusa/cad-http/src/handler"
	"github.com/joshbrusa/cad-http/src/middleware"
)

func (server *Server) StartServer() {
	middleware := middleware.NewMiddleware(server.Logger)
	handler := handler.NewHandler(server.Logger)

	server.Mux.HandleFunc("/", middleware.ApplyMiddleware(
		handler.HandleRoot(),
		middleware.LoggerMiddleware,
	))

	fmt.Println("server listening on port:", server.Port)
	err := http.ListenAndServe(":"+server.Port, server.Mux)

	if err != nil {
		server.Logger.Error(err.Error())
		os.Exit(1)
	}
}
