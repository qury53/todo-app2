package tasks_repository

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, taskID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	DELETE
	FROM todoapp.tasks
	WHERE id = $1;
	`

	cmdTag, err := r.Pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("task not found: %w", core_errors.ErrNotFound)
	}

	return nil
}
