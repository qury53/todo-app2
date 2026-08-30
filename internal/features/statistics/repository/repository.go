package statistics_repository

import core_pgx_pool "github.com/qury53/todo-app2/internal/core/repository/postgres"

type StatisticsRepository struct {
	Pool *core_pgx_pool.Pool
}

func NewStatisticsRepository(pool *core_pgx_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{
		Pool: pool,
	}
}
