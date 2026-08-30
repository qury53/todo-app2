package users_repository

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (r *UsersRepository) CreateUser(ctx context.Context, userDomain domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO todoapp.users (full_name, password, phone_number, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id, version, full_name, password, phone_number, email;
	`

	fmt.Println(userDomain)
	row := r.Pool.QueryRow(
		ctx,
		query,
		userDomain.FullName,
		userDomain.Password,
		userDomain.PhoneNumber,
		userDomain.Email,
	)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.Password,
		&userModel.PhoneNumber,
		&userModel.Email,
	); err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	user := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.Password,
		userModel.PhoneNumber,
		userModel.Email,
	)

	fmt.Println(*user.PhoneNumber)
	return user, nil
}
