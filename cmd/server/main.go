package main

import (
	"github.com/joshbrusa/cad-core/internal/logger"
	"github.com/joshbrusa/cad-core/internal/server"
)

func main() {
	logger := logger.New()

	server := server.New(logger)
	server.Start()
}
