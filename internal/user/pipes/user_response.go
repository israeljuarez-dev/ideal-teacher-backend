package pipes

import (
	"time"

	"github.com/google/uuid"
	user "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
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

	GetUserByEmailOut struct {
		ID        uuid.UUID
		Email     string
		FirstName string
		LastName  string
		Status    user.UserStatus
		CreatedAt time.Time
		RoleName  string
	}
)
