package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	routerAuth "github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/routes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	routerUser "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/routes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

const (
	pathPrefix = "/api/v1"
)

func InitRouters(db *postgres.DB, v *validator.Validator, cfg *config.JWT) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares útiles
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routerUser.SetUpUser(pathPrefix, r, db)
	routerAuth.SetUpAuth(pathPrefix, r, db, v, cfg)

	return r
}
