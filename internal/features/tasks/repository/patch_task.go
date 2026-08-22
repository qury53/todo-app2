package tasks_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/todo-app2/internal/core/domain"
)

func (r *TasksRepository) PatchTask(
	ctx context.Context,
	taskID int,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()
	query := `
	UPDATE todoapp.tasks
	SET
	    title = $1,
	    description = $2,
	    completed = $3,
    	completed_at = $4,
    	version = version + 1
	WHERE id = $5 AND version = $6
	RETURNING
	    id, version, title, description, completed, created_at, completed_at, author_id;
`
	row := r.Pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		taskID,
		task.Version,
	)

	var taskModel TaskModel

	if err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorID,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("task is accessed concurrently: %w", err)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return domainFromModel(taskModel), nil
}
