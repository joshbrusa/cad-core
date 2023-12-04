package server

import (
	"net/http"
	"os"

	"github.com/joshbrusa/cad-http/src/logger"
	"github.com/joshbrusa/cad-http/src/postgres"
)

type Server struct {
	Logger   *logger.Logger
	Mux      *http.ServeMux
	Postgres *postgres.Postgres
	Port     string
}

func NewServer(logger *logger.Logger, postgres *postgres.Postgres) *Server {
	return &Server{
		Logger:   logger,
		Mux:      http.NewServeMux(),
		Postgres: postgres,
		Port:     os.Getenv("PORT"),
	}
}
