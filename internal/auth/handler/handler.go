package handler

import (
	"log/slog"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

type (
	AuthHandler interface {
		Login(w http.ResponseWriter, r *http.Request)
	}

	Handler struct {
		serv service.AuthService
		log  *slog.Logger
		v    *validator.Validator
	}
)

func New(serv service.AuthService, log *slog.Logger, v *validator.Validator) *Handler {
	return &Handler{
		serv: serv,
		log:  log,
		v:    v,
	}
}
