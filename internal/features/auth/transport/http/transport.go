package auth_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/qury53/todo-app2/internal/core/domain"
	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
)

type AuthTransport struct {
	AuthService AuthService
}

type AuthService interface {
	IsLoginAndPasswordValid(
		ctx context.Context,
		username string,
		password string,
	) (int, bool, error)
	GenerateAccessToken(userID int) (string, error)
	GenerateRefreshToken(userID int) (string, error)
	GetRefreshTokenTTL() time.Duration
	CreateUser(ctx context.Context, userDomain domain.User) (domain.User, error)
	ValidateRefreshToken(ctx context.Context, tokenStr string) (int, bool, error)
}

func NewAuthTransport(service AuthService) *AuthTransport {
	return &AuthTransport{
		AuthService: service,
	}
}

func (t *AuthTransport) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/auth/login/",
			Handler: t.Login,
			Method:  http.MethodPost,
		},
		{
			Path:    "/auth/register/",
			Handler: t.Register,
			Method:  http.MethodPost,
		},
		{
			Path:    "/auth/refresh/",
			Handler: t.Refresh,
			Method:  http.MethodPost,
		},
	}
}
