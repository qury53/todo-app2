package tasks_repository

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (r *TasksRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.tasks (
	title, description, completed, created_at, completed_at, author_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, version, title, description, completed, created_at, completed_at, author_id;
	`

	if task.Completed {
		task.CompletedAt = &task.CreatedAt
	}

	row := r.Pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorID,
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
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return domainFromModel(taskModel), nil
}
