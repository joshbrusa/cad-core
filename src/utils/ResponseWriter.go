package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseWriter struct {
	Logger *slog.Logger
}

func NewResponseWriter(logger *slog.Logger) *ResponseWriter {
	return &ResponseWriter{
		Logger: logger,
	}
}

func (responseWriter *ResponseWriter) WriteCode(
	w http.ResponseWriter,
	code int,
) {
	w.WriteHeader(code)
}

func (responseWriter *ResponseWriter) WriteJson(
	w http.ResponseWriter,
	code int,
	data map[string]any,
) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseWriter.Logger.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
