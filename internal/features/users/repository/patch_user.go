package users_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	userID int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	fmt.Println(user.Version)
	query := `
	UPDATE todoapp.users
	SET
		full_name=$1,
		password=$2,
		phone_number=$3,
		email=$4,
		version=version+1
	WHERE id=$5 AND version=$6
	RETURNING
		id,
		version,
		full_name,
		password,
		phone_number,
		email
	`

	row := r.Pool.QueryRow(
		ctx,
		query,
		user.FullName,
		user.Password,
		user.PhoneNumber,
		user.Email,
		userID,
		user.Version,
	)

	var patchedModel UserModel

	if err := row.Scan(
		&patchedModel.ID,
		&patchedModel.Version,
		&patchedModel.FullName,
		&patchedModel.Password,
		&patchedModel.PhoneNumber,
		&patchedModel.Email,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user is concurrently accessed: %w", core_errors.ErrConflict)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	fmt.Println(patchedModel.Version)
	return domainFromModel(patchedModel), nil
}
