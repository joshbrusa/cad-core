package middleware

import (
	"net/http"
)

func (middleware *Middleware) LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.Logger.Info(
			"logging middleware",
			"path", r.URL.Path,
			"method", r.Method,
		)
		next.ServeHTTP(w, r)
	}
}
