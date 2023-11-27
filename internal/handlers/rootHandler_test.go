package handlers

import (
	"testing"

	"github.com/joshbrusa/cad-core/internal/loggers"
)

func TestRootHandler(t *testing.T) {
	jsonLogger := loggers.NewJsonLogger()
	rootHandler := NewRootHandler(jsonLogger)

	rootHandler.Handle()
}
