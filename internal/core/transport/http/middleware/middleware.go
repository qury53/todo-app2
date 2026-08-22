package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	mLen := len(m)
	if mLen == 0 {
		return h
	}
	for i := mLen - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}
