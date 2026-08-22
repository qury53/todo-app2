package transport_http_response

import (
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     domain.StatusCodeUninitialized,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *ResponseWriter) GetStatusCode() int {
	return w.statusCode
}
