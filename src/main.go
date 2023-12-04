package main

import "github.com/joshbrusa/cad-core/src/core"

func main() {
	logger := core.NewLogger()
	postgres := core.NewPostgres(logger)
	server := core.NewServer(logger, postgres)
	server.Start()
}
