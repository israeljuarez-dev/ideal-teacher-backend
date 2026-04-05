package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	authRoutes "github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/routes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/routing"
	userRoutes "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/routes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

const (
	pathPrefix = "/api/v1"
)

func InitRouters(db *postgres.DB, v *validator.Validator, cfg *config.JWT) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	groups := []routing.Group{
		userRoutes.InitUserRoutes(db, cfg),
		authRoutes.InitAuthRoutes(db, v, cfg),
	}

	r.Route(pathPrefix, func(r chi.Router) {
		RegisterRoutes(r, groups)
	})

	return r
}
