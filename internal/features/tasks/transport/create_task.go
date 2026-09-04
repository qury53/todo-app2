package tasks_transport_http

import (
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=100" example:"go walking"`
	Description string `json:"description" validate:"required,max=1000" example:"walk 10k steps"`
	Completed   bool   `json:"completed" example:"false"`
	AuthorID    int    `json:"author_id" validate:"required" example:"10"`
}

type CreateTaskResponse TaskDTO

// @Summary creates task
// Description creates task in db
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body CreateTaskRequest true "CreateTask request body"
// @Success 201 {object} CreateTaskResponse "Successfully created task"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /task/ [post]
func (t *TasksTransport) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.RespondError("failed to create task", err)
		return
	}

	task := domainFromDTO(request)
	task, err := t.TasksService.CreateTask(ctx, task)
	if err != nil {
		responseHandler.RespondError("failed to create task", err)
		return
	}

	responseHandler.RespondJSON(http.StatusCreated, CreateTaskResponse(task))
}

func domainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(
		dto.Title,
		dto.Description,
		dto.Completed,
		dto.AuthorID,
	)
}
