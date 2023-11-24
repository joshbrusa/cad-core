package database

import (
	"os"

	"github.com/jackc/pgx"
)

func New() *Database {
	config := pgx.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	db, err := pgx.Connect(config)

	if err != nil {
		os.Exit(1)
	}

	return &Database{
		DB: db,
	}
}
