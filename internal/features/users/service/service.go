package users_service

import (
	"context"

	"github.com/qury53/todo-app2/internal/core/domain"
)

type UserService struct {
	UserRepository UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, userDomain domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, userID int) (domain.User, error)
	PatchUser(ctx context.Context, userID int, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
}

func NewUsersService(UserRepository UserRepository) *UserService {
	return &UserService{
		UserRepository: UserRepository,
	}
}
