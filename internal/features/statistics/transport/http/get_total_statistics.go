package statistics_transport

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

type GetTotalStatisticsResponse TotalStatisticsDTO

func (t *StatisticsTransport) GetTotalStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)
	stats, err := t.StatisticsService.GetTotalStatistics(ctx)
	if err != nil {
		responseHandler.RespondError("failed to get statistics", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetTotalStatisticsResponse(dtoFromDomainTotal(stats)))
}
