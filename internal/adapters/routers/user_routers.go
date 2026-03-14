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

	api.HandleFunc(userIDPath, userHandler.GetByIDHandler).Methods("GET")
	api.HandleFunc(userBasicPath, userHandler.GetByEmailHandler).Methods("GET")

	api.HandleFunc(userBasicPath, userHandler.RegisterHandler).Methods("POST")

	api.HandleFunc(usersPath, userHandler.GetAllHandler).Methods("GET")

	api.HandleFunc(userIDPath, userHandler.UpdateHandler).Methods("PUT")
	
	api.HandleFunc(userIDPath, userHandler.DeleteHandler).Methods("DELETE")
}
