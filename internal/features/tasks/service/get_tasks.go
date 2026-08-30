package tasks_service

import (
	"context"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	tasks, err := s.TasksRepository.GetTasks(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
