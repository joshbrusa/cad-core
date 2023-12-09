package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/utils"
)

type PanicMiddleware struct {
	Logger *core.Logger
}

func NewPanicMiddleware(logger *core.Logger) *PanicMiddleware {
	return &PanicMiddleware{
		Logger: logger,
	}
}

func (panicMiddleware *PanicMiddleware) Handle(next http.Handler) http.Handler {
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
			if ok {
				slog.Error(recMsg)
			}

			// send response
			response := utils.NewResponse(w, http.StatusInternalServerError, msg)
			response.Send()
		}()
		next.ServeHTTP(w, r)
	})
}
