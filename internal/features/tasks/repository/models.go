package tasks_repository

import (
	"time"

	"github.com/qury53/todo-app2/internal/core/domain"
)

type TaskModel struct {
	ID          int        `db:"id"`
	Version     int        `db:"version"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	Completed   bool       `db:"completed"`
	CreatedAt   time.Time  `db:"created_at"`
	CompletedAt *time.Time `db:"completed_at"`
	AuthorID    int        `db:"author_id"`
}

func domainFromModel(model TaskModel) domain.Task {
	return domain.Task{
		ID:          model.ID,
		Version:     model.Version,
		Title:       model.Title,
		Description: model.Description,
		Completed:   model.Completed,
		CreatedAt:   model.CreatedAt,
		CompletedAt: model.CompletedAt,
		AuthorID:    model.AuthorID,
	}
}

func domainsFromModels(models []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))

	for i, model := range models {
		domains[i] = domainFromModel(model)
	}

	return domains

}
