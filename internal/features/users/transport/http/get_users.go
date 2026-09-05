package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTO

// GetUsers godoc
// @Summary Gets users
// Description Gets users from db
// @Tags users
// @Produce json
// @Success 200 {object} GetUsersResponse "Successfully got users"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/ [get]
func (t *UserTransport) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)

	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	logger.Info("getting query params")
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.RespondError("failed to get limit/offset params", err)
		return
	}
	logger.Debug("getting users from service")
	users, err := t.UserService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.RespondError("failed to get users", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, dtosFromDomains(users))
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get limit param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get offset param: %w", err)
	}

	return limit, offset, nil
}
