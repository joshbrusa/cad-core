package core

import (
	"errors"
	"os"
)

type Env struct {
	Port             string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
}

func NewEnv(logger *Logger) (*Env, error) {
	port, portOk := os.LookupEnv("PORT")

	if !portOk {
		return nil, errors.New("missing environment variable: PORT")
	}

	postgresHost, postgresHostOk := os.LookupEnv("POSTGRES_HOST")

	if !postgresHostOk {
		return nil, errors.New("missing environment variable: POSTGRES_HOST")
	}

	postgresUser, postgresUserOk := os.LookupEnv("POSTGRES_USER")

	if !postgresUserOk {
		return nil, errors.New("missing environment variable: POSTGRES_USER")
	}

	postgresPassword, postgresPasswordOk := os.LookupEnv("POSTGRES_PASSWORD")

	if !postgresPasswordOk {
		return nil, errors.New("missing environment variable: POSTGRES_PASSWORD")
	}

	postgresDb, postgresDbOk := os.LookupEnv("POSTGRES_DB")

	if !postgresDbOk {
		return nil, errors.New("missing environment variable: POSTGRES_DB")
	}

	return &Env{
		Port:             port,
		PostgresHost:     postgresHost,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDb:       postgresDb,
	}, nil
}
