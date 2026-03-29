package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/role/domain"
)

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatysBlocked  UserStatus = "blocked"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string
	Status    UserStatus
	CreatedAt time.Time
	Roles     domain.Role
}

type Users []User
