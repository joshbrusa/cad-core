package server

import (
	"net/http"

	"github.com/joshbrusa/cad-core/internal/logger"
)

type Server struct {
	Logger *logger.Logger
	Mux    *http.ServeMux
	Port   string
}
