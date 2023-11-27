package main

import (
	"github.com/joshbrusa/cad-core/internal/databases"
	"github.com/joshbrusa/cad-core/internal/loggers"
	"github.com/joshbrusa/cad-core/internal/servers"
)

func main() {
	jsonLogger := loggers.NewJsonLogger()

	postgres := databases.NewPostgres(jsonLogger)
	httpServer := servers.NewHttpServer(jsonLogger, postgres)

	httpServer.Start()
}
