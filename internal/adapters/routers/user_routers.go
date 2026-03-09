package routers

import (
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/handler"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/storage/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/storage/postgres/repository"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/core/service"
	"github.com/gorilla/mux"

)

const (
	userBasicPath = "/user"
	userIDPath    = "/user/{id}"
	usersPath     = "/users"
)

func setUpUser(router *mux.Router, db *postgres.DB) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api := router.PathPrefix(apiV1).Subrouter()

	api.HandleFunc(userIDPath, userHandler.GetByID).Methods("GET")
	api.HandleFunc(userBasicPath, userHandler.GetByEmail).Methods("GET")

	api.HandleFunc(userBasicPath, userHandler.Register).Methods("POST")

	api.HandleFunc(usersPath, userHandler.GetAll).Methods("GET")

	api.HandleFunc(userIDPath, userHandler.Update).Methods("PUT")
	
	api.HandleFunc(userIDPath, userHandler.Delete).Methods("DELETE")
}
