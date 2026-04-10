package models

import (
	"time"

	"github.com/google/uuid"
	user "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
)

type (
	GetUserByEmailOut struct {
		ID        uuid.UUID
		Email     string
		FirstName string
		LastName  string
		Password  string
		Status    user.UserStatus
		CreatedAt time.Time
		RoleName  string
	}
)
