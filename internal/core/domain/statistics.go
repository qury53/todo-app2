package domain

import (
	"time"
)

type UserStatistics struct {
	TotalTasks          int
	CompletedTasks      int
	CompletedPrecentage float64
	AvgComletionTime    time.Duration
}

type TotalStatistics struct {
	TotalUsers          int
	TotalTasks          int
	CompletedTasks      int
	CompletedPrecentage float64
	AvgComletionTime    time.Duration
}

func ProcessUserStatistics(tasks []Task) UserStatistics {
	var stats UserStatistics
	total := len(tasks)
	stats.TotalTasks = total

	var totalTime time.Duration
	var completed int
	for _, task := range tasks {
		if task.Completed {
			completed++
			totalTime += task.CompletedAt.Sub(task.CreatedAt)
		}
	}

	stats.CompletedTasks = completed
	stats.CompletedPrecentage = (float64(completed) / float64(total)) * 100
	stats.AvgComletionTime = totalTime / time.Duration(completed)

	return stats
}

func ProcessTotalStatistics(tasks []Task, users []User) TotalStatistics {
	var stats TotalStatistics
	totalTasks := len(tasks)
	totalUsers := len(users)

	stats.TotalTasks = totalTasks
	stats.TotalUsers = totalUsers

	var totalTime time.Duration
	var completed int
	for _, task := range tasks {
		if task.Completed {
			completed++
			totalTime += task.CompletedAt.Sub(task.CreatedAt)
		}
	}

	stats.CompletedTasks = completed
	stats.CompletedPrecentage = (float64(completed) / float64(totalTasks)) * 100
	stats.AvgComletionTime = totalTime / time.Duration(completed)

	return stats
}
