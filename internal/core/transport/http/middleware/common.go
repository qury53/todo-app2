package core_http_middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(logger *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := logger.With(
				"request_id", requestID,
			)

			ctx := core_logger.WithLogger(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					logger := core_logger.FromContext(r.Context())
					responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)
					responseHandler.RespondPanic(p, "got unexprected panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := core_logger.FromContext(r.Context())
			rw := transport_http_response.NewResponseWriter(w)

			logger.Debug("...incoming HTTP request")
			t := time.Now()

			next.ServeHTTP(rw, r)

			logger.Debug(
				"finnished handling HTTP request",
				slog.String("method", r.Method),
				slog.Int("status", rw.GetStatusCode()),
				slog.Duration("time spent", time.Since(t)),
			)
		})
	}
}
