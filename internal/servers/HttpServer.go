package servers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshbrusa/cad-core/internal/databases"
	"github.com/joshbrusa/cad-core/internal/handlers"
	"github.com/joshbrusa/cad-core/internal/loggers"
)

type HttpServer struct {
	JsonLogger *loggers.JsonLogger
	Mux        *http.ServeMux
	Database   *databases.Postgres
	Port       string
}

func NewHttpServer(JsonLogger *loggers.JsonLogger, database *databases.Postgres) *HttpServer {
	return &HttpServer{
		JsonLogger: JsonLogger,
		Mux:        http.NewServeMux(),
		Database:   database,
		Port:       os.Getenv("PORT"),
	}
}

func (httpServer *HttpServer) Start() {
	rootHandler := handlers.NewRootHandler(httpServer.JsonLogger)

	httpServer.Mux.Handle("/", rootHandler.Handler())

	fmt.Println("server listening on port:", httpServer.Port)
	err := http.ListenAndServe(":"+httpServer.Port, httpServer.Mux)

	if err != nil {
		httpServer.JsonLogger.Error(err.Error())
		os.Exit(1)
	}
}
