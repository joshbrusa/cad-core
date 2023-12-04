package middleware

import "github.com/joshbrusa/cad-http/src/logger"

type Middleware struct {
	Logger *logger.Logger
}

func NewMiddleware(logger *logger.Logger) *Middleware {
	return &Middleware{
		Logger: logger,
	}
}
