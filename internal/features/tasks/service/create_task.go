package tasks_service

import (
	"context"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *TasksService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, err
	}

	task, err := s.TasksRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}
