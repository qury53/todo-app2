package statistics_repository

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
	tasks_repository "github.com/qury53/todo-app2/internal/features/tasks/repository"
)

func (r *StatisticsRepository) GetTasksByUserID(
	ctx context.Context,
	userID int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_id
	FROM todoapp.tasks
	WHERE id = $1;
	`

	var tasksModels []tasks_repository.TaskModel
	rows, err := r.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	for rows.Next() {
		var taskModel tasks_repository.TaskModel
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
			return nil, fmt.Errorf("scan error: %w", err)
		}

		tasksModels = append(tasksModels, taskModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return domainsFromModels(tasksModels), nil
}
