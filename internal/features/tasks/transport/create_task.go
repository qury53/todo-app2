package tasks_transport_http

import (
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,max=1000"`
	Completed   bool   `json:"completed"`
	AuthorID    int    `json:"author_id" validate:"required"`
}

type CreateTaskResponse TaskDTO

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
