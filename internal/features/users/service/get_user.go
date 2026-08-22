package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *UserService) GetUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := s.UserRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
