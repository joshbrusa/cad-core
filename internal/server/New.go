package server

import (
	"net/http"
	"os"

	"github.com/joshbrusa/cad-core/internal/database"
	"github.com/joshbrusa/cad-core/internal/logger"
)

func New(slog *logger.Slog, database *database.Database) *Server {
	return &Server{
		Slog:     slog,
		Mux:      http.NewServeMux(),
		Database: database,
		Port:     os.Getenv("PORT"),
	}
}
