package statistics_repository

import (
	"github.com/qury53/todo-app2/internal/core/domain"
	tasks_repository "github.com/qury53/todo-app2/internal/features/tasks/repository"
	users_repository "github.com/qury53/todo-app2/internal/features/users/repository"
)

func domainFromModel(model tasks_repository.TaskModel) domain.Task {
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

func domainsFromModels(models []tasks_repository.TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))

	for i, model := range models {
		domains[i] = domainFromModel(model)
	}

	return domains

}

func userDomainFromModel(model users_repository.UserModel) domain.User {
	return domain.User{
		ID:          model.ID,
		Version:     model.Version,
		FullName:    model.FullName,
		Password:    model.Password,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
	}
}

func userDomainsFromModels(models []users_repository.UserModel) []domain.User {
	domains := make([]domain.User, len(models))
	for i, model := range models {
		domains[i] = userDomainFromModel(model)
	}

	return domains
}
