package auth_service

import (
	"context"
	"errors"
	"fmt"

	core_encrypt "github.com/qury53/todo-app2/internal/core/encrypt"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (s *AuthService) IsLoginAndPasswordValid(
	ctx context.Context,
	username string,
	password string,
) (int, bool, error) {
	user, err := s.AuthRepository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return -1, false, core_errors.ErrInvalidCredentials
		}

		return -1, false, fmt.Errorf("failed to get user: %w", err)
	}

	if err := core_encrypt.CompareHashAndPassword(
		user.Password,
		password,
	); err != nil {
		return user.ID, false, core_errors.ErrInvalidCredentials
	}

	return user.ID, true, nil
}
