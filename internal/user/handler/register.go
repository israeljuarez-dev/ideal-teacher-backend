package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/myerrors"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/pipes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/pkg/response"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req pipes.CreateUserIn

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("register: invalid JSON body", "error", err)
		response.Error(w, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	if fieldErrs := h.v.Validate(req); fieldErrs != nil {
		h.log.Error("register: validation failed", "fields", len(fieldErrs))
		response.ValidationError(w, fieldErrs)
		return
	}

	u := &service.InsertUserIn{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	userResp, err := h.serv.Register(r.Context(), u)
	if err != nil {
		var userErr *myerrors.UserError
		if errors.As(err, &userErr) {
			switch {
			case errors.Is(userErr.Err, myerrors.EmailAlreadyExistsError):
				h.log.Error("register: email already exists", "email", req.Email)
				response.Error(w, response.ErrorResponse{
					Status:  http.StatusConflict,
					Message: userErr.Msg,
				})
			default:
				h.log.Error("register: domain error", "error", userErr)
				response.Error(w, response.ErrorResponse{
					Status:  http.StatusUnprocessableEntity,
					Message: userErr.Msg,
				})
			}
			return
		}

		h.log.Error("register: unexpected error", "error", err)
		response.Error(w, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	h.log.Info("register: user created", "user_id", userResp.ID)
	response.JSON(w, response.Success[*service.InsertUserOut]{
		Status:  http.StatusCreated,
		Message: "user registered successfully",
		Data:    userResp,
	})
}
