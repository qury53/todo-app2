package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

func (t *TasksTransport) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	taskID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("failed to get task id", err)
		return
	}

	if err := t.TasksService.DeleteTask(ctx, taskID); err != nil {
		responseHandler.RespondError("failed to delete task", err)
		return
	}

	responseHandler.RespondNoContent()
}
