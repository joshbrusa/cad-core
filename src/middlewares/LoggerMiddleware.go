package middlewares

import (
	"log/slog"
	"net/http"
)

type LoggerMiddleware struct {
	Logger *slog.Logger
}

func NewLoggerMiddleware(logger *slog.Logger) *LoggerMiddleware {
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
