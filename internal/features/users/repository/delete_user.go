package users_repository

import (
	"context"
	"fmt"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	DELETE
	FROM todoapp.users
	WHERE id = $1
	`

	cmdTag, err := r.Pool.Exec(ctx, query, userID)

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found: %w", core_errors.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("exec delete error: %w", err)
	}

	return nil
}
