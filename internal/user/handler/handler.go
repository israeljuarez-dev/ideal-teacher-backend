package handler

import (
	"log/slog"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

type (
	UserHandler interface {
		Register(w http.ResponseWriter, r *http.Request)
	}

	Handler struct {
		serv service.UserService
		log  *slog.Logger
		v    *validator.Validator
	}
)

func New(serv service.UserService, log *slog.Logger, v *validator.Validator) *Handler {
	return &Handler{
		serv: serv,
		log:  log,
		v:    v,
	}
}
