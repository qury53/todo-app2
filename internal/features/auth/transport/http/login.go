package auth_transport_http

import (
	"net/http"
	"time"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=100" example:"sadamhussein2014"`
	Password string `json:"password" validate:"required,min=6,max=100" example:"qwerty123"`
}

type LoginAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// Login godoc
// @Summary Login user
// Description checks if username and password are valid and returns tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginUserRequest true "LoginUser request body"
// @Success 200 {object} LoginAccessTokenResponse "Successfully logged in user"
// @Failure 401 {object} transport_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /auth/login/ [post]
func (t *AuthTransport) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	var loginUserRequest LoginUserRequest
	if err := core_http_request.DecodeAndValidate(
		r,
		&loginUserRequest,
	); err != nil {
		responseHandler.RespondError("failed to decode and validate request", err)
		return
	}

	userID, valid, err := t.AuthService.IsLoginAndPasswordValid(
		ctx,
		loginUserRequest.Username,
		loginUserRequest.Password,
	)

	if err != nil {
		responseHandler.RespondError("failed to check if login/password is correct", err)
		return
	}

	if !valid {
		responseHandler.RespondStatusCode(http.StatusUnauthorized)
		return
	}

	accessToken, err := t.AuthService.GenerateAccessToken(userID)
	if err != nil {
		responseHandler.RespondError("failed to generate access token", err)
		return
	}

	refreshToken, err := t.AuthService.GenerateRefreshToken(userID)
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
		LoginAccessTokenResponse{
			AccessToken: accessToken,
			TokenType:   "Bearer",
		},
	)
}
