package auth_transport_http

import (
	"net/http"
	"time"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// Refresh godoc
// @Summary Refresh access token
// @Description Uses refresh cookie to issue a new access token and rotate the refresh cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} RefreshAccessTokenResponse "Successfully refreshed tokens"
// @Failure 401 {object} transport_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /auth/refresh/ [post]
func (t *AuthTransport) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	// 1. Извлекаем refresh_token из куки
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		// Куки нет — пользователь не авторизован
		responseHandler.RespondStatusCode(http.StatusUnauthorized)
		return
	}

	// 2. Валидируем рефреш-токен через сервис (он должен распарсить JWT и вернуть userID)
	userID, valid, err := t.AuthService.ValidateRefreshToken(ctx, cookie.Value)
	if err != nil {
		responseHandler.RespondError("failed to validate refresh token", err)
		return
	}

	if !valid {
		responseHandler.RespondStatusCode(http.StatusUnauthorized)
		return
	}

	// 3. Генерируем новую пару токенов (Token Rotation)
	newAccessToken, err := t.AuthService.GenerateAccessToken(userID)
	if err != nil {
		responseHandler.RespondError("failed to generate new access token", err)
		return
	}

	newRefreshToken, err := t.AuthService.GenerateRefreshToken(userID)
	if err != nil {
		responseHandler.RespondError("failed to generate new refresh token", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(t.AuthService.GetRefreshTokenTTL() / time.Second),
	})

	// 5. Отдаем новый Access Token в JSON
	responseHandler.RespondJSON(
		http.StatusOK,
		RefreshAccessTokenResponse{
			AccessToken: newAccessToken,
			TokenType:   "Bearer",
		},
	)
}
