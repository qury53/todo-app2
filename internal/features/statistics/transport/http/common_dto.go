package statistics_transport

import (
	"github.com/qury53/todo-app2/internal/core/domain"
)

type UserStatisticsDTO struct {
	TotalTasks          int     `json:"total_tasks"`
	CompletedTasks      int     `json:"completed_tasks"`
	CompletedPrecentage float64 `json:"completed_precentage"`
	AvgComletionTime    string  `json:"avg_completion_time"`
}

type TotalStatisticsDTO struct {
	TotalUsers          int     `json:"total_users"`
	TotalTasks          int     `json:"total_tasks"`
	CompletedTasks      int     `json:"completed_tasks"`
	CompletedPrecentage float64 `json:"completed_precentage"`
	AvgComletionTime    string  `json:"avg_completion_time"`
}

func dtoFromDomainUser(stats domain.UserStatistics) UserStatisticsDTO {
	return UserStatisticsDTO{
		TotalTasks:          stats.TotalTasks,
		CompletedTasks:      stats.CompletedTasks,
		CompletedPrecentage: stats.CompletedPrecentage,
		AvgComletionTime:    stats.AvgComletionTime.String(),
	}
}

func dtoFromDomainTotal(stats domain.TotalStatistics) TotalStatisticsDTO {
	return TotalStatisticsDTO{
		TotalUsers:          stats.TotalUsers,
		TotalTasks:          stats.TotalTasks,
		CompletedTasks:      stats.CompletedTasks,
		CompletedPrecentage: stats.CompletedPrecentage,
		AvgComletionTime:    stats.AvgComletionTime.String(),
	}
}
