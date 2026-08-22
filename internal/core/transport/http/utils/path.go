package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	val := r.PathValue(key)
	if val == "" {
		return 0, fmt.Errorf("%s is not set: %w", key, core_errors.ErrInvalidArgument)
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf(
			"%s has to be valid integer: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	return intVal, nil
}
