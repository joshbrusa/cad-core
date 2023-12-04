package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(logger *Logger) *Postgres {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err1 := sql.Open("postgres", config)

	if err1 != nil {
		logger.Error(err1.Error())
		os.Exit(1)
	}

	err2 := db.Ping()

	if err2 != nil {
		logger.Error(err2.Error())
		os.Exit(1)
	}

	return &Postgres{
		DB: db,
	}
}
