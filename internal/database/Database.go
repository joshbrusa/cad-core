package database

import "github.com/jackc/pgx"

type Database struct {
	DB *pgx.Conn
}
