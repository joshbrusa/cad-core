package core

import "os"

type Env struct {
	Port             string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
}

func NewEnv(logger *Logger) *Env {
	port, portOk := os.LookupEnv("PORT")

	if !portOk {
		logger.Error("missing environment variable: PORT")
		os.Exit(1)
	}

	postgresHost, postgresHostOk := os.LookupEnv("POSTGRES_HOST")

	if !postgresHostOk {
		logger.Error("missing environment variable: POSTGRES_HOST")
		os.Exit(1)
	}

	postgresUser, postgresUserOk := os.LookupEnv("POSTGRES_USER")

	if !postgresUserOk {
		logger.Error("missing environment variable: POSTGRES_USER")
		os.Exit(1)
	}

	postgresPassword, postgresPasswordOk := os.LookupEnv("POSTGRES_PASSWORD")

	if !postgresPasswordOk {
		logger.Error("missing environment variable: POSTGRES_PASSWORD")
		os.Exit(1)
	}

	postgresDb, postgresDbOk := os.LookupEnv("POSTGRES_DB")

	if !postgresDbOk {
		logger.Error("missing environment variable: POSTGRES_DB")
		os.Exit(1)
	}

	return &Env{
		Port:             port,
		PostgresHost:     postgresHost,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDb:       postgresDb,
	}
}
