package auth_transport_http

import (
	"net/http"
	"time"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	users_transport_http "github.com/qury53/todo-app2/internal/features/users/transport/http"
)

type CreateUserResponse users_transport_http.UserDTO

type CreateUserRequest struct {
	Username    string  `json:"username" validate:"required,min=3,max=100" example:"qerfasd33"`
	FullName    string  `json:"full_name" validate:"required,min=3,max=100" example:"Ivan Ivanov"`
	Password    string  `json:"password" validate:"required,min=6,max=100" example:"qwerty123"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164" example:"+79998887766"`
	Email       string  `json:"email" validate:"required,email" example:"somemail@mail.ru"`
}

type RegisterUserResponse struct {
	LoginAccessTokenResponse
	User CreateUserResponse `json:"user"`
}

// Register godoc
// @Summary creates user
// Description creates user in db
// @Tags auth
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "CreateUser request body"
// @Success 201 {object} RegisterUserResponse "Successfully registered user"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /auth/register/ [post]
func (t *AuthTransport) Register(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest
	ctx := r.Context()

	logger := core_logger.FromContext(ctx)
	logger.Debug("started creating user")

	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	if err := core_http_request.DecodeAndValidate(r, &createUserRequest); err != nil {
		responseHandler.RespondError("validate request", err)
		return
	}

	userDomain := domainFromDTO(createUserRequest)
	userDomain, err := t.AuthService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.RespondError("failed to create user", err)
		return
	}

	accessToken, err := t.AuthService.GenerateAccessToken(userDomain.ID)
	if err != nil {
		responseHandler.RespondError("failed to generate access token", err)
		return
	}

	refreshToken, err := t.AuthService.GenerateRefreshToken(userDomain.ID)
	if err != nil {
		responseHandler.RespondError("failed to generate refresh token", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(t.AuthService.GetRefreshTokenTTL() / time.Second),
	})

	responseHandler.RespondJSON(
		http.StatusOK,
		RegisterUserResponse{
			LoginAccessTokenResponse: LoginAccessTokenResponse{
				AccessToken: accessToken,
				TokenType:   "Bearer",
			},
			User: CreateUserResponse(dtoFromDomain(userDomain)),
		},
	)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(
		dto.Username,
		dto.FullName,
		dto.Password,
		dto.PhoneNumber,
		dto.Email,
	)
}

func dtoFromDomain(domain domain.User) users_transport_http.UserDTO {
	return users_transport_http.UserDTO{
		ID:          domain.ID,
		Version:     domain.Version,
		Username:    domain.Username,
		FullName:    domain.FullName,
		Password:    domain.Password,
		PhoneNumber: domain.PhoneNumber,
		Email:       domain.Email,
	}
}
