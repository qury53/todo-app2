package tasks_transport_http

import (
	"time"

	"github.com/qury53/todo-app2/internal/core/domain"
)

type TaskDTO struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	AuthorID    int        `json:"author_id"`
}

func dtoFromDomain(taskDomain domain.Task) TaskDTO {
	return TaskDTO{
		ID:          taskDomain.ID,
		Version:     taskDomain.Version,
		Title:       taskDomain.Title,
		Description: taskDomain.Description,
		Completed:   taskDomain.Completed,
		CreatedAt:   taskDomain.CreatedAt,
		CompletedAt: taskDomain.CompletedAt,
		AuthorID:    taskDomain.AuthorID,
	}
}

func dtosFromDomains(taskDomains []domain.Task) []TaskDTO {
	dtos := make([]TaskDTO, len(taskDomains))
	for i, domain := range taskDomains {
		dtos[i] = dtoFromDomain(domain)
	}

	return dtos
}
