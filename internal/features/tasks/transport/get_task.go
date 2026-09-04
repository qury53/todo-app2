package tasks_transport_http

import (
	"net/http"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

type GetTaskResponse TaskDTO

// GetTask godoc
// @Summary Gets task
// Description Gets task from db
// @Tags tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} GetTaskResponse "Successfully got task"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [get]
func (t *TasksTransport) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	taskID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError(
			"failed to get task id: %w",
			core_errors.ErrInvalidArgument,
		)

		return
	}

	task, err := t.TasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.RespondError("failed to get task", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetTaskResponse(task))
}
