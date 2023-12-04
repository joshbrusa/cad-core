package middleware

import (
	"log/slog"
	"net/http"

	"github.com/joshbrusa/cad-http/src/response"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rec *http.Request) {
		defer func() {
			rec := recover()

			if rec == nil {
				return
			}

			msg := "Internal server error."

			slog.Error(msg)

			response := response.NewResponse(w, http.StatusInternalServerError, msg)

			response.SendResponse()

			recMsg, ok := rec.(string)

			if !ok {
				slog.Error("Wrong panic type.")
				return
			}

			slog.Error(recMsg)
		}()

		next.ServeHTTP(w, rec)
	})
}
