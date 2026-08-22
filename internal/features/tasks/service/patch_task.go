package tasks_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	taskID int,
	patch domain.TaskPatch,
) (domain.Task, error) {
	task, err := s.TasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	task, err = s.TasksRepository.PatchTask(ctx, taskID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to patch task: %w", err)
	}

	fmt.Println(task)

	return task, nil
}
