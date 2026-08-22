package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

// GetIntQueryParam() returns query param of type integer. If param doesn't
// exists returns nil value and nil error. If value is not valid integer returns
// nil value and error ErrInvalidArgument. If value is valid returns it.
func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	v := r.URL.Query().Get(key)
	if v == "" {
		return nil, nil
	}

	value, err := strconv.Atoi(v)
	if err != nil {
		return nil, fmt.Errorf(
			"query param is not valid integer: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return &value, nil
}
