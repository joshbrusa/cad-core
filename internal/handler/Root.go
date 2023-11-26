package handler

import "net/http"

func (handler *Handler) Root() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.Slog.Info("handling root", "path", r.URL.Path)

		w.WriteHeader(http.StatusOK)
	})
}
