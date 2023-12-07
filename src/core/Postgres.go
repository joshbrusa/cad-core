package core

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(env *Env, logger *Logger) (*Postgres, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		env.PostgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDb,
	)

	db, dbErr := sql.Open("postgres", config)

	if dbErr != nil {
		return nil, dbErr
	}

	pingErr := db.Ping()

	if pingErr != nil {
		return nil, dbErr
	}

	return &Postgres{
		DB: db,
	}, nil
}
