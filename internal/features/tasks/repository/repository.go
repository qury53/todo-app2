package tasks_repository

import core_pgx_pool "github.com/qury53/todo-app2/internal/core/repository/postgres"

type TasksRepository struct {
	Pool *core_pgx_pool.Pool
}

func NewTasksRepository(pool *core_pgx_pool.Pool) *TasksRepository {
	return &TasksRepository{
		Pool: pool,
	}
}
