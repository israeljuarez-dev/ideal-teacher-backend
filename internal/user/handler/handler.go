package handler

import (
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
)

type (
	userHandler interface {
		Register(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		serv service.UserService
	}
)

func New(serv service.UserService) userHandler {
	return &handler{serv: serv}
}
