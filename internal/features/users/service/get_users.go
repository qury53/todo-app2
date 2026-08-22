package users_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	users, err := s.UserRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		fmt.Println("repo error")
		return nil, err
	}

	return users, nil
}
