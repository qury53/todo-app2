package auth_repostitory

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
	users_repository "github.com/qury53/todo-app2/internal/features/users/repository"
)

func (r *AuthRepository) GetUserByUsername(
	ctx context.Context,
	username string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, username, full_name, password, phone_number, email
	FROM todoapp.users
	WHERE username=$1
	`

	row := r.Pool.QueryRow(ctx, query, username)
	var userModel users_repository.UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Username,
		&userModel.FullName,
		&userModel.Password,
		&userModel.PhoneNumber,
		&userModel.Email,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user not found: %w", core_errors.ErrNotFound)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return domainFromModel(userModel), nil
}

func domainFromModel(model users_repository.UserModel) domain.User {
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
