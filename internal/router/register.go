package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/routing"
)

func RegisterRoutes(r chi.Router, groups []routing.Group) {
	for _, group := range groups {
		r.Group(func(r chi.Router) {
			for _, mdw := range group.Middlewares {
				r.Use(mdw)
			}

			r.Route(group.Prefix, func(r chi.Router) {
				for _, route := range group.Routes {
					r.With(route.Middlewares...).MethodFunc(
						route.Method,
						route.Path,
						route.Handler,
					)
				}
			})
		})
	}
}
