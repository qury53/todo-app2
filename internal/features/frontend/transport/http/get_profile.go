package frontend_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

func (t *FrontendTransport) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	file, err := t.FrontendService.GetProfile()
	if err != nil {
		responseHandler.RespondError("failed to get profile", err)
		return
	}

	responseHandler.RespondFile(http.StatusOK, file)
}
