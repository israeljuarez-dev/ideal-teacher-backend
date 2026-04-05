package handler

import (
	"encoding/json"
	"net/http"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/pipes"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/service"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req pipes.CreateUserIn

	defer r.Body.Close()
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResp)
}
