package routes

import (
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/handler"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/routing"
	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

const (
	basePath      = "/auth"
	authLoginPath = "/login"
)

func InitAuthRoutes(db *postgres.DB, v *validator.Validator, cfg *config.JWT) routing.Group {
	repo := repository.New(db)
	serv := service.New(repo, cfg)
	hand := handler.New(serv, v)

	return routing.Group{
		Routes: []routing.Route{
			{
				Method:  http.MethodPost,
				Path:    basePath + authLoginPath,
				Handler: hand.Login,
			},
		},
	}
}
