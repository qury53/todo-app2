package auth_repostitory

import core_pgx_pool "github.com/qury53/todo-app2/internal/core/repository/postgres"

type AuthRepository struct {
	Pool *core_pgx_pool.Pool
}

func NewAuthRepository(pool *core_pgx_pool.Pool) *AuthRepository {
	return &AuthRepository{
		Pool: pool,
	}
}
