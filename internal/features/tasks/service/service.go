package tasks_service

import (
	"context"

	"github.com/qury53/todo-app2/internal/core/domain"
)

type TasksService struct {
	TasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	GetTasks(ctx context.Context, limit *int, offset *int) ([]domain.Task, error)
	PatchTask(ctx context.Context, taskID int, task domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
}

func NewTasksService(repository TasksRepository) *TasksService {
	return &TasksService{
		TasksRepository: repository,
	}
}
