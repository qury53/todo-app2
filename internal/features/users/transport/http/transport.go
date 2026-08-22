package users_transport_http

import (
	"context"
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
)

type UserTransport struct {
	UserService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, userDomain domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, userID int) (domain.User, error)
	PatchUser(ctx context.Context, userID int, patch domain.UserPatch) (domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
}

func NewUserTransport(service UserService) *UserTransport {
	return &UserTransport{
		UserService: service,
	}
}

func (t *UserTransport) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users/",
			Handler: t.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/",
			Handler: t.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: t.GetUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: t.PatchUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: t.DeleteUser,
		},
	}
}
