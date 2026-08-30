package tasks_service

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (s *TasksService) DeleteTask(ctx context.Context, taskID int) error {
	if taskID < 1 {
		return fmt.Errorf("taskID has to be positive: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.TasksRepository.DeleteTask(ctx, taskID); err != nil {
		return err
	}

	return nil
}
