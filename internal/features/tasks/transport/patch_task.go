package tasks_transport_http

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_types "github.com/qury53/todo-app2/internal/core/transport/http/types"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title" example:"some_title"`
	Description core_http_types.Nullable[string] `json:"description" example:"some_description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed" example:"true"`
}

type PatchTaskResponse TaskDTO

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		titleLen := utf8.RuneCountInString(*r.Title.Value)
		if titleLen < 1 && titleLen > 100 {
			return fmt.Errorf(
				"title has to be between 1 and 100: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Description.Set {
		descriptionLen := utf8.RuneCountInString(*r.Description.Value)
		if descriptionLen > 1000 {
			return fmt.Errorf(
				"description length has to not more than 1000: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf(
				"completed cannot be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

// PatchTask godoc
// @Summary Patches task
// Description Patches task in db
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param patch body PatchTaskRequest true "PatchRequestBody"
// @Success 200 {object} PatchTaskResponse "Patched task response"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} transport_http_response.ErrorResponse "Conflict"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /tasks/{id} [patch]
func (t *TasksTransport) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	taskID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError(
			"failed to get task id",
			core_errors.ErrInvalidArgument,
		)

		return
	}

	var patchRequest PatchTaskRequest

	if err := core_http_request.DecodeAndValidate(r, &patchRequest); err != nil {
		responseHandler.RespondError("failed to decode request", err)
		return
	}

	if err := patchRequest.Validate(); err != nil {
		responseHandler.RespondError("failed to validate request", err)
		return
	}

	taskPatch := taskPatchFromRequest(patchRequest)
	task, err := t.TasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.RespondError("failed to patch task", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, PatchTaskResponse(task))
}

func taskPatchFromRequest(patchRequest PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		patchRequest.Title.ToDomain(),
		patchRequest.Description.ToDomain(),
		patchRequest.Completed.ToDomain(),
	)
}
