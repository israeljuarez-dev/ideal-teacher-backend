package handler

import (
	"encoding/json"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/pipes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/service"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req pipes.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.v.Validate(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validator.FormatValidationErrors(err))
		return
	}

	l := &service.LoginIn{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := h.serv.Login(r.Context(), l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
