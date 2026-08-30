package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
)

type TasksTransport struct {
	TasksService TasksService
}

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	GetTasks(ctx context.Context, limit *int, offset *int) ([]domain.Task, error)
	PatchTask(ctx context.Context, taskID int, patch domain.TaskPatch) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
}

func NewTasksTransport(tasksService TasksService) *TasksTransport {
	return &TasksTransport{
		TasksService: tasksService,
	}
}
func (t *TasksTransport) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Path:    "/tasks/",
			Method:  http.MethodPost,
			Handler: t.CreateTask,
		},
		{
			Path:    "/tasks/{id}",
			Method:  http.MethodGet,
			Handler: t.GetTask,
		},
		{
			Path:    "/tasks/{id}",
			Method:  http.MethodDelete,
			Handler: t.DeleteTask,
		},
		{
			Path:    "/tasks/{id}",
			Method:  http.MethodPatch,
			Handler: t.PatchTask,
		},
		{
			Path:    "/tasks/",
			Method:  http.MethodGet,
			Handler: t.GetTasks,
		},
	}
}
