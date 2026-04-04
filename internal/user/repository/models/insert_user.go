package models

import (
	"time"

	"github.com/google/uuid"
	user "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	InsertUserParamsIn struct {
		ID        pgtype.UUID
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
)
