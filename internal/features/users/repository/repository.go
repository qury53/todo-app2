package users_repository

import core_pgx_pool "github.com/qury53/todo-app2/internal/core/repository/postgres"

type UsersRepository struct {
	Pool *core_pgx_pool.Pool
}

func NewUsersRepository(pool *core_pgx_pool.Pool) *UsersRepository {
	return &UsersRepository{
		Pool: pool,
	}
}
