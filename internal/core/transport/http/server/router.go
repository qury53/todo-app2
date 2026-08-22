package transport_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/qury53/todo-app2/internal/core/transport/http/middleware"
)

type APIVersion string

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion  APIVersion
	middlewares []core_http_middleware.Middleware
}

func NewAPIVersionRouter(
	apiVersion APIVersion,
	m ...core_http_middleware.Middleware,
) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:    http.NewServeMux(),
		apiVersion:  apiVersion,
		middlewares: m,
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(r, r.middlewares...)
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		h := route.WithMiddleware()
		r.Handle(pattern, h)
	}
}
