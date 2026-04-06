package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/myerrors"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/pipes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/pkg/response"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req pipes.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("login: invalid JSON body", "error", err)
		response.Error(w, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	if fieldErrs := h.v.Validate(req); fieldErrs != nil {
		h.log.Warn("login: validation failed", "fields", len(fieldErrs))
		response.ValidationError(w, fieldErrs)
		return
	}

	l := &service.LoginIn{
		Email:    req.Email,
		Password: req.Password,
	}

	out, err := h.serv.Login(r.Context(), l)
	if err != nil {
		var authErr *myerrors.AuthError
		if errors.As(err, &authErr) {
			switch {
			case errors.Is(authErr.Err, myerrors.InvalidEmailOrPassword):
				h.log.Warn("login: user not found", "email", req.Email)
			default:
				h.log.Warn("login: auth error", "error", authErr.Msg)
			}
			response.Error(w, response.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: authErr.Msg,
			})
			return
		}

		h.log.Error("login: unexpected error", "error", err)
		response.Error(w, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	h.log.Info("login: success", "email", req.Email)
	response.JSON(w, response.Success[*service.LoginOut]{
		Status:  http.StatusOK,
		Message: "login successful",
		Data:    out,
	})
}
