package middlewares

import (
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
)

type LoggerMiddleware struct {
	Logger *core.Logger
}

func NewLoggerMiddleware(logger *core.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		Logger: logger,
	}
}

func (loggerMiddleware *LoggerMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggerMiddleware.Logger.Info(
			"request",
			"path", r.URL.Path,
			"method", r.Method,
		)
		next.ServeHTTP(w, r)
	})
}
