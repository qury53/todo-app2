package frontend_transport_http

import (
	"net/http"

	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
)

type FrontendTransport struct {
	FrontendService FrontendService
}

type FrontendService interface {
	GetMainPage() ([]byte, error)
	GetProfile() ([]byte, error)
}

func NewFrontendTransport(service FrontendService) *FrontendTransport {
	return &FrontendTransport{
		FrontendService: service,
	}
}

func (t *FrontendTransport) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/",
			Handler: t.GetMainPage,
		},
		{
			Path:    "/profile/",
			Handler: t.GetProfile,
			Method:  http.MethodGet,
		},
	}
}
