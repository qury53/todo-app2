package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *UserService) PatchUser(
	ctx context.Context,
	userID int,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := s.UserRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	if err = user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("failed to apply patch to user: %w", err)
	}

	user, err = s.UserRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
