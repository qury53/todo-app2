package users_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (r *UsersRepository) GetUser(ctx context.Context, userID int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, username, full_name, password, phone_number, email
	FROM todoapp.users
	WHERE id = $1;
	`

	var user UserModel

	row := r.Pool.QueryRow(ctx, query, userID)
	err := row.Scan(
		&user.ID,
		&user.Version,
		&user.Username,
		&user.FullName,
		&user.Password,
		&user.PhoneNumber,
		&user.Email,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user not found: %w", core_errors.ErrNotFound)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return domainFromModel(user), nil
}
