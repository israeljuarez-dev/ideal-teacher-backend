package handler

import (
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
)

type (
	UserHandler interface {
		Register(w http.ResponseWriter, r *http.Request)
	}

	Handler struct {
		serv service.UserService
	}
)

func New(serv service.UserService) *Handler {
	return &Handler{serv: serv}
}
