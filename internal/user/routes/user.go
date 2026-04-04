package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/handler"
	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
)

const (
	basePath = "/users"
	userIDPath    = "/users/{id}"
)

func SetUpUser(pathPrefix string, router chi.Router, db *postgres.DB) {
	userRepository := repository.New(db)
	userService := service.New(userRepository)
	userHandler := handler.New(userService)

	router.Route(basePath, func(r chi.Router) {
		r.Post(basePath, userHandler.Register)
	})
}
