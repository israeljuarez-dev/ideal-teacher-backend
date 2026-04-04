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

	handler struct {
		serv service.AuthService
		v    *validator.Validator
	}
)

func New(serv service.AuthService, v *validator.Validator) AuthHandler {
	return &handler{
		serv: serv,
		v:    v,
	}
}
