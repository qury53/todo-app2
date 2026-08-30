package tasks_repository

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_id
	FROM todoapp.tasks
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var tasksModels []TaskModel

	for rows.Next() {
		var taskModel TaskModel
		if err := rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorID,
		); err != nil {
			return nil, fmt.Errorf("scan error", err)
		}

		tasksModels = append(tasksModels, taskModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return domainsFromModels(tasksModels), nil
}
