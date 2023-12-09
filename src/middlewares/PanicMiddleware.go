package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/response"
)

// handles an uncaught panic
func PanicMiddleware(logger *core.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				// get recover, if its nil there is no panic
				rec := recover()
				if rec == nil {
					return
				}

				// log internal server error
				msg := "internal server error"
				slog.Error(msg)

				// log message if string
				recMsg, ok := rec.(string)
				if !ok {
					slog.Error("panic not a string")
				} else {
					slog.Error(recMsg)
				}

				// send response
				response := response.NewResponse(w, http.StatusInternalServerError, msg)
				response.SendResponse()
			}()
			next.ServeHTTP(w, r)
		})
	}
}
