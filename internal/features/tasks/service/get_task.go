package tasks_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (s *TasksService) GetTask(ctx context.Context, taskID int) (domain.Task, error) {
	if taskID < 1 {
		return domain.Task{}, fmt.Errorf(
			"taskID has to be positive number: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	task, err := s.TasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}
