package utils

import "net/http"

// Middleware takes an HTTP handler and a variable number of middleware functions.
// It applies each middleware to the handler in the order they are provided.
// This allows chaining multiple middleware functions in a flexible manner.
func Middleware(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
