package utils

import (
	"encoding/json"
	"net/http"

	"github.com/joshbrusa/cad-http/src/core"
)

type ResponseWriter struct {
	Logger *core.Logger
}

func NewResponseWriter(logger *core.Logger) *ResponseWriter {
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
