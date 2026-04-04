package pipes

import (
	"github.com/google/uuid"
)

type (
	UserResponseOut struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		FullName string    `json:"full_name"`
		Role     string    `json:"role"`
		Status   string    `json:"status"`
	}

	UsersListResponseOut struct {
		Data  []UserResponseOut `json:"data"`
		Total int               `json:"total"`
	}
)
