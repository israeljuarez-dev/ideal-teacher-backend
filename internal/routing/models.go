package routing

import "net/http"

var (
	AppRouting = []Group{}
)

type (
	Group struct {
		Prefix      string
		Middlewares []func(http.Handler) http.Handler
		Routes      []Route
	}

	Route struct {
		Method      string
		Path        string
		Handler     http.HandlerFunc
		Middlewares []func(http.Handler) http.Handler
	}
)
