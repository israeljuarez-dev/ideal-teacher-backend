package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
)

const (
	apiV1 = "/api/v1"
)

func InitRouters(db *postgres.DB) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares útiles
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	setUpUser(r, db)

	return r
}
