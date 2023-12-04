package main

import (
	"github.com/joshbrusa/cad-http/src/logger"
	"github.com/joshbrusa/cad-http/src/postgres"
	"github.com/joshbrusa/cad-http/src/server"
)

func main() {
	logger := logger.NewLogger()
	postgres := postgres.NewPostgres(logger)
	server := server.NewServer(logger, postgres)
	server.StartServer()
}
