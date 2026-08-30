package statistics_transport

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

type GetUserStatisticsResponse UserStatisticsDTO

func (t *StatisticsTransport) GetUserStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("failed to get user id", err)
		return
	}

	stats, err := t.StatisticsService.GetUserStatistics(ctx, userID)
	if err != nil {
		responseHandler.RespondError("failed to get user statistics", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, GetUserStatisticsResponse(dtoFromDomainUser(stats)))
}
