package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/handler"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/repository"
)

const (
	userBasicPath = "/users"
	userIDPath    = "/users/{id}"
)

func setUpUser(router chi.Router, db *postgres.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.Route(apiV1, func(r chi.Router) {
		r.Post(userBasicPath, userHandler.Register)
	})
}
