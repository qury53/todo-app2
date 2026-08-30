package statistics_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *StatisticsService) GetTotalStatistics(ctx context.Context) (domain.TotalStatistics, error) {
	tasks, err := s.StatisticsRepository.GetTasks(ctx)
	if err != nil {
		return domain.TotalStatistics{}, fmt.Errorf("failed to get tasks: %w", err)
	}

	users, err := s.StatisticsRepository.GetUsers(ctx)
	if err != nil {
		return domain.TotalStatistics{}, fmt.Errorf("failed to get users: %w", err)
	}

	return domain.ProcessTotalStatistics(tasks, users), nil
}
