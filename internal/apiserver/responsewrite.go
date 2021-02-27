package apiserver

import (
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.Header().Add(HeaderContentType, "application/json")
	w.ResponseWriter.WriteHeader(statusCode)
}
