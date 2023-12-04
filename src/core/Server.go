package core

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshbrusa/cad-core/src/handlers"
)

type Server struct {
	Logger   *Logger
	Mux      *http.ServeMux
	Postgres *Postgres
	Port     string
}

func NewServer(logger *Logger, postgres *Postgres) *Server {
	return &Server{
		Logger:   logger,
		Mux:      http.NewServeMux(),
		Postgres: postgres,
		Port:     os.Getenv("PORT"),
	}
}

func (server *Server) Start() {
	rootHandler := handlers.NewRootHandler(server.Logger)

	server.Mux.Handle("/", rootHandler.Handle())

	fmt.Println("server listening on port:", server.Port)
	err := http.ListenAndServe(":"+server.Port, server.Mux)

	if err != nil {
		server.Logger.Error(err.Error())
		os.Exit(1)
	}
}
