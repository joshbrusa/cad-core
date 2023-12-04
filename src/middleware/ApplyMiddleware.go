package middleware

import "net/http"

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func (middleware *Middleware) ApplyMiddleware(handlerFunc http.HandlerFunc, middlewareFunc ...MiddlewareFunc) http.HandlerFunc {
	for _, middleware := range middlewareFunc {
		handlerFunc = middleware(handlerFunc)
	}
	return handlerFunc
}
