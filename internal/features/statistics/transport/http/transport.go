package statistics_transport

import (
	"context"
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
	transport_http_server "github.com/qury53/todo-app2/internal/core/transport/http/server"
)

type StatisticsTransport struct {
	StatisticsService StatisticsService
}

type StatisticsService interface {
	GetUserStatistics(ctx context.Context, userID int) (domain.UserStatistics, error)
	GetTotalStatistics(ctx context.Context) (domain.TotalStatistics, error)
}

func NewStatisticsTransport(statisticsService StatisticsService) *StatisticsTransport {
	return &StatisticsTransport{
		StatisticsService: statisticsService,
	}
}

func (t *StatisticsTransport) Routes() []transport_http_server.Route {
	return []transport_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics/{id}",
			Handler: t.GetUserStatistics,
		},
		{
			Method:  http.MethodGet,
			Path:    "/statistics/",
			Handler: t.GetTotalStatistics,
		},
	}
}
