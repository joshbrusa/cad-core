package server

import (
	"net/http"

	"github.com/joshbrusa/cad-core/internal/database"
	"github.com/joshbrusa/cad-core/internal/logger"
)

type Server struct {
	Slog     *logger.Slog
	Mux      *http.ServeMux
	Database *database.Database
	Port     string
}
