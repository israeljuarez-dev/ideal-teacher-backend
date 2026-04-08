package routes

import (
	"log/slog"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/routing"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/handler"
	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
	"github.com/israeljuarez-dev/ideal-teacher-backend/pkg/middleware"
)

const (
	basePath   = "/users"
	userIDPath = "/users/{id}"
	voidPath   = "/"
)

func InitUserRoutes(db *postgres.DB, v *validator.Validator, log *slog.Logger, cfg *config.JWT) routing.Group {
	repo := repository.New(db)
	serv := service.New(repo, log)
	hand := handler.New(serv, log, v)

	return routing.Group{
		Prefix: basePath,
		Routes: []routing.Route{
			{
				Method:  http.MethodPost,
				Path:    voidPath,
				Handler: hand.Register,
			},
			{
				Method:  http.MethodGet,
				Path:    voidPath,
				Handler: hand.GetByEmail,
				Middlewares: []func(http.Handler) http.Handler{
					middleware.ValidateJWT(cfg, log),
				},
			},
		},
	}
}
