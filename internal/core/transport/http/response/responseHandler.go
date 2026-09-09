package transport_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
)

type HTTPResponseHandler struct {
	logger *core_logger.Logger
	w      http.ResponseWriter
}

func NewHTTPResponseHandler(
	logger *core_logger.Logger,
	w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		logger: logger,
		w:      w,
	}
}

func (h *HTTPResponseHandler) RespondNoContent() {
	h.w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) RespondJSON(statusCode int, data any) {
	h.w.WriteHeader(statusCode)
	if err := json.NewEncoder(h.w).Encode(data); err != nil {
		h.logger.Error("failed to write HTTP response")
	}
}

func (h *HTTPResponseHandler) RespondFile(statusCode int, file []byte) {
	h.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.w.WriteHeader(statusCode)
	_, err := h.w.Write(file)
	if err != nil {
		h.logger.Error("failed to write file", "error", err)
	}
}

func (h *HTTPResponseHandler) RespondStatusCode(statusCode int) {
	h.w.WriteHeader(statusCode)
}

func (h *HTTPResponseHandler) RespondError(message string, err error) {
	var (
		statusCode int
		logFunc    func(msg string, args ...any)
	)

	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.logger.Info
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.logger.Info
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.logger.Warn
	case errors.Is(err, core_errors.ErrInvalidCredentials):
		statusCode = http.StatusUnauthorized
		logFunc = h.logger.Debug
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.logger.Error
	}

	logFunc(message, "error", err)
	h.errorRespond(statusCode, message, err)
}

func (h *HTTPResponseHandler) errorRespond(statusCode int, message string, err error) {
	h.w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:   err.Error(),
		Message: message,
	}

	if err := json.NewEncoder(h.w).Encode(response); err != nil {
		h.logger.Error("failed to write err HTTP response")
	}
}

func (h *HTTPResponseHandler) RespondPanic(p any, message string) {
	h.logger.Error("got panic", "panic", p)
	h.errorRespond(
		http.StatusInternalServerError,
		message,
		fmt.Errorf("internal server error"),
	)
}
