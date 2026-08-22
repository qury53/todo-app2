package transport_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_middleware "github.com/qury53/todo-app2/internal/core/transport/http/middleware"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *Config
	logger *core_logger.Logger
	m      []core_http_middleware.Middleware
}

func NewHTTPServer(config *Config, logger *core_logger.Logger, m ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
		logger: logger,
		m:      m,
	}
}

func (s *HTTPServer) RegisterRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		r := router.WithMiddleware()
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, r))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.m...)
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTime)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("failed to shutdown server gracefully: %w", err)
		}
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("run error: %w", err)
		}
	}

	return nil
}
