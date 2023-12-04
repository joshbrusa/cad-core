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

func NewPostgres(logger *Logger, env *Env) *Postgres {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		env.PostgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDb,
	)

	db, dbErr := sql.Open("postgres", config)

	if dbErr != nil {
		logger.Error(dbErr.Error())
		os.Exit(1)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		logger.Error(pingErr.Error())
		os.Exit(1)
	}

	return &Postgres{
		DB: db,
	}
}
