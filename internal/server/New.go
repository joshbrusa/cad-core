package server

import (
	"net/http"
	"os"

	"github.com/joshbrusa/cad-core/internal/logger"
)

func New(logger *logger.Logger) *Server {
	return &Server{
		Logger: logger,
		Mux:    http.NewServeMux(),
		Port:   os.Getenv("PORT"),
	}
}
