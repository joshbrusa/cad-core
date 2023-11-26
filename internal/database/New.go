package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joshbrusa/cad-core/internal/logger"
	_ "github.com/lib/pq"
)

func New(slog *logger.Slog) *Database {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err1 := sql.Open("postgres", config)

	if err1 != nil {
		slog.Error(err1.Error())
		os.Exit(1)
	}

	err2 := db.Ping()

	if err2 != nil {
		slog.Error(err2.Error())
		os.Exit(1)
	}

	return &Database{
		DB: db,
	}
}
