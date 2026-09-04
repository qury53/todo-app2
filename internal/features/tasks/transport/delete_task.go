package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

// DeleteTask godoc
// @Summary Delete task
// Description delete task from db
// @Tags tasks
// @Accept json
// @Param id path int true "Task ID"
// @Success 204 "No content"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [delete]
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
