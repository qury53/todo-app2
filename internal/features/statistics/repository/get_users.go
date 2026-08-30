package statistics_repository

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
	users_repository "github.com/qury53/todo-app2/internal/features/users/repository"
)

func (r *StatisticsRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, password, phone_number, email
	FROM todoapp.users
	`

	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	var users []users_repository.UserModel

	for rows.Next() {
		var user users_repository.UserModel
		rows.Scan(
			&user.ID,
			&user.Version,
			&user.FullName,
			&user.Password,
			&user.PhoneNumber,
			&user.Email,
		)

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		users = append(users, user)
	}

	return userDomainsFromModels(users), nil
}
