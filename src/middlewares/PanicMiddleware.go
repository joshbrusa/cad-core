package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
	"github.com/joshbrusa/cad-http/src/utils"
)

type PanicMiddleware struct {
	Logger         *core.Logger
	ResponseWriter *utils.ResponseWriter
}

func NewPanicMiddleware(
	logger *core.Logger,
	responseWriter *utils.ResponseWriter,
) *PanicMiddleware {
	return &PanicMiddleware{
		Logger:         logger,
		ResponseWriter: responseWriter,
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

			// log message
			recMsg, ok := rec.(string)

			if ok {
				slog.Error(recMsg)
			} else {
				slog.Error("internal server error")
			}

			// send response
			panicMiddleware.ResponseWriter.WriteCode(w, http.StatusInternalServerError)
		}()
		next.ServeHTTP(w, r)
	})
}
