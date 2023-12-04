package main

import (
	"github.com/joshbrusa/cad-core/src/logger"
	"github.com/joshbrusa/cad-core/src/postgres"
	"github.com/joshbrusa/cad-core/src/server"
)

func main() {
	logger := logger.NewLogger()
	postgres := postgres.NewPostgres(logger)
	server := server.NewServer(logger, postgres)
	server.Start()
}
