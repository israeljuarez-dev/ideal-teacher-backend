package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/handler"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/repository"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

const (
	authBasicPath = "/auth/login"
)

func setUpAuth(router chi.Router, db *postgres.DB, v *validator.Validator, cfg *config.JWT) {
	authRepository := repository.NewUserRepository(db)
	authServ := service.NewAuthService(authRepository, cfg)
	authHandler := handler.NewAuthHandler(authServ, v)

	router.Route(apiV1, func(r chi.Router) {
		r.Post(authBasicPath, authHandler.Login)
	})
}
