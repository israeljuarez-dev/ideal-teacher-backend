package dto

import (
	"github.com/google/uuid"
)

type (
	UserResponse struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		FullName string    `json:"full_name"`
		Role     string    `json:"role"`
		Status   string    `json:"status"`
	}

	UsersListResponse struct {
		Data  []UserResponse `json:"data"`
		Total int            `json:"total"`
	}
)
