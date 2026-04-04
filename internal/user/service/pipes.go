package service

import (
	"time"

	"github.com/google/uuid"
	user "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
)

type (
	InsertUserIn struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
	}

	InsertUserOut struct {
		ID        uuid.UUID
		Email     string
		FirstName string
		LastName  string
		Status    user.UserStatus
		CreatedAt time.Time
	}

	GetByEmailUserOut struct {
		ID        uuid.UUID
		Email     string
		FirstName string
		LastName  string
		Status    user.UserStatus
		CreatedAt time.Time
		RoleName  string
	}
)
