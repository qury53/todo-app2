package users_repository

import "github.com/qury53/todo-app2/internal/core/domain"

type UserModel struct {
	ID          int     `db:"id"`
	Version     int     `db:"version"`
	Username    string  `db:"username"`
	FullName    string  `db:"full_name"`
	Password    string  `db:"password"`
	PhoneNumber *string `db:"phone_number"`
	Email       string  `db:"email"`
}

func domainFromModel(model UserModel) domain.User {
	return domain.User{
		ID:          model.ID,
		Version:     model.Version,
		Username:    model.Username,
		FullName:    model.FullName,
		Password:    model.Password,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
	}
}

func domainsFromModels(models []UserModel) []domain.User {
	domains := make([]domain.User, len(models))
	for i, model := range models {
		domains[i] = domainFromModel(model)
	}

	return domains
}
