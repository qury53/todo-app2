package frontend_transport_http

import (
	"net/http"

	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

func (t *FrontendTransport) GetMainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	file, err := t.FrontendService.GetMainPage()
	if err != nil {
		responseHandler.RespondError("failed to return index.html", err)
		return
	}

	responseHandler.RespondFile(http.StatusOK, file)
}
