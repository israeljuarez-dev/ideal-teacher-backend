package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/handler"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

const (
	authBasicPath = "/auth/login"
)

func SetUpAuth(pathPrefix string, router chi.Router, db *postgres.DB, v *validator.Validator, cfg *config.JWT) {
	authRepository := repository.New(db)
	authServ := service.New(authRepository, cfg)
	authHandler := handler.New(authServ, v)

	router.Route(pathPrefix, func(r chi.Router) {
		r.Post(authBasicPath, authHandler.Login)
	})
}
