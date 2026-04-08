package handler

import (
	"errors"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/myerrors"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/pipes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/pkg/response"
)

func (h *Handler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	var req pipes.GetUserByEmailIn

	req.Email = r.URL.Query().Get("email")
	if fieldErrs := h.v.Validate(req); fieldErrs != nil {
		h.log.Error("get_by_email: validation failed", "fields", len(fieldErrs))
		response.ValidationError(w, fieldErrs)
		return
	}

	userResp, err := h.serv.GetByEmail(r.Context(), req.Email)
	if err != nil {
		var userErr *myerrors.UserError
		if errors.As(err, &userErr) {
			switch {
			case errors.Is(userErr.Err, myerrors.UserNotFoundError):
				h.log.Error("get_by_email: user not found", "email", req.Email)
				response.Error(w, response.ErrorResponse{
					Status:  http.StatusNotFound,
					Message: userErr.Msg,
				})
			default:
				h.log.Error("get_by_email: domain error", "error", userErr)
				response.Error(w, response.ErrorResponse{
					Status:  http.StatusUnprocessableEntity,
					Message: userErr.Msg,
				})
			}
			return
		}

		h.log.Error("get_by_email: unexpected error", "error", err)
		response.Error(w, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	h.log.Info("get_by_email: user found", "email", req.Email)
	response.JSON(w, response.Success[*pipes.GetUserByEmailOut]{
		Status:  http.StatusOK,
		Message: "user retrieved successfully",
		Data: &pipes.GetUserByEmailOut{
			ID:        userResp.ID,
			Email:     userResp.Email,
			FirstName: userResp.FirstName,
			LastName:  userResp.LastName,
			Status:    userResp.Status,
			CreatedAt: userResp.CreatedAt,
			RoleName:  userResp.RoleName,
		},
	})
}
