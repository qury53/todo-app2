package statistics_service

import (
	"context"
	"fmt"

	"github.com/qury53/todo-app2/internal/core/domain"
)

func (s *StatisticsService) GetUserStatistics(
	ctx context.Context,
	userID int,
) (domain.UserStatistics, error) {
	tasks, err := s.StatisticsRepository.GetTasksByUserID(ctx, userID)
	if err != nil {
		return domain.UserStatistics{}, fmt.Errorf("failed to get tasks by user id: %w", err)
	}

	stats := domain.ProcessUserStatistics(tasks)
	return stats, nil
}
