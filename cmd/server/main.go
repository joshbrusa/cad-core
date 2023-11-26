package main

import (
	"github.com/joshbrusa/cad-core/internal/database"
	"github.com/joshbrusa/cad-core/internal/logger"
	"github.com/joshbrusa/cad-core/internal/server"
)

func main() {
	logger := logger.New()

	database := database.New(logger)
	server := server.New(logger, database)

	server.Start()
}
