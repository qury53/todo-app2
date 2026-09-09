package users_repository

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()
	fmt.Println("inside repo")

	query := `
	SELECT id, version, username, full_name, password, phone_number, email
	FROM todoapp.users
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	fmt.Println("inside repo 2")
	var users []UserModel

	for rows.Next() {
		var user UserModel
		rows.Scan(
			&user.ID,
			&user.Version,
			&user.Username,
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

	return domainsFromModels(users), nil
}
