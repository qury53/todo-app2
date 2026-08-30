package statistics_service

import (
	"context"

	"github.com/qury53/todo-app2/internal/core/domain"
)

type StatisticsService struct {
	StatisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetTasksByUserID(ctx context.Context, userID int) ([]domain.Task, error)
	GetTasks(ctx context.Context) ([]domain.Task, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
}

func NewStatisticsService(statisticsRepository StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		StatisticsRepository: statisticsRepository,
	}
}
