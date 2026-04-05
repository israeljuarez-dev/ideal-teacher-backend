package handler

import (
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
		v    *validator.Validator
	}
)

func New(serv service.AuthService, v *validator.Validator) *Handler {
	return &Handler{
		serv: serv,
		v:    v,
	}
}
