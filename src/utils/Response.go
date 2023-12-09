package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	w          http.ResponseWriter
	StatusCode int
	Data       any
}

func (response *Response) Send() {
	jsonData, _ := json.Marshal(response.Data)

	response.w.Header().Set("Content-Type", "application/json")
	response.w.WriteHeader(response.StatusCode)
	response.w.Write(jsonData)
}
