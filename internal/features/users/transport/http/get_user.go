package users_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

type GetUserDTO UserDTO

func (t *UserTransport) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("failed to get user id", err)
		return
	}

	user, err := t.UserService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.RespondError("failed to get user", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetUserDTO(dtoFromDomain(user)))
}

// func getUserID(r *http.Request) (int, error) {
// 	val := r.PathValue("id")
// 	if val == "" {
// 		return 0, fmt.Errorf("user id is not set: %w", core_errors.ErrInvalidArgument)
// 	}
//
// 	userID, err := strconv.Atoi(val)
// 	if err != nil {
// 		return 0, fmt.Errorf(
// 			"user id has to be valid integer: %w",
// 			core_errors.ErrInvalidArgument,
// 		)
// 	}
//
// 	return userID, nil
// }
