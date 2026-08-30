package tasks_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (r *TasksRepository) GetTask(ctx context.Context, taskID int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_id
	FROM todoapp.tasks
	WHERE id = $1
	`

	row := r.Pool.QueryRow(ctx, query, taskID)
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
			return domain.Task{}, fmt.Errorf("task not found: %w", core_errors.ErrNotFound)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return domainFromModel(taskModel), nil
}
