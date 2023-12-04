package response

import "net/http"

type Response struct {
	ResponseWriter http.ResponseWriter
	StatusCode     int
	Data           interface{}
}

func NewResponse(responseWriter http.ResponseWriter, statusCode int, data interface{}) *Response {
	return &Response{
		ResponseWriter: responseWriter,
		StatusCode:     statusCode,
		Data:           data,
	}
}
