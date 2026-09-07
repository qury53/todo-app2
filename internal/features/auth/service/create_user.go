package auth_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_encrypt "github.com/qury53/todo-app2/internal/core/encrypt"
)

func (s *AuthService) CreateUser(ctx context.Context, userDomain domain.User) (domain.User, error) {
	if err := userDomain.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("create user service: %w", err)
	}

	hashedPassword, err := core_encrypt.HashPassword(userDomain.Password)

	if err != nil {
		return domain.User{}, fmt.Errorf("service error", err)
	}

	userDomain.Password = hashedPassword
	fmt.Println(userDomain)
	userDomain, err = s.AuthRepository.CreateUser(ctx, userDomain)

	if err != nil {
		return domain.User{}, fmt.Errorf("create user service: %w", err)
	}
	fmt.Println(userDomain)
	return userDomain, nil
}
