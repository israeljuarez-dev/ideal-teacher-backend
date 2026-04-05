package routes

import (
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/routing"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/handler"
	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
	// "github.com/israeljuarez-dev/ideal-teacher-backend/pkg/middleware"
)

const (
	basePath   = "/users"
	userIDPath = "/users/{id}"
)

func InitUserRoutes(db *postgres.DB, cfg *config.JWT) routing.Group {
	repo := repository.New(db)
	serv := service.New(repo)
	hand := handler.New(serv)

	return routing.Group{
		Routes: []routing.Route{
			{
				Method:  http.MethodPost,
				Path:    basePath,
				Handler: hand.Register,
				/*Middlewares: []func(http.Handler) http.Handler{
					middleware.ValidateJWT(cfg),
				},*/
			},
		},
	}
}
