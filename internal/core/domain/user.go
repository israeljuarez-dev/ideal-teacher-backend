package domain

import (
	"time"

)

type Role string

type Status string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
	RoleAdmin   Role = "admin"

	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

type User struct {
	ID        int32
	Email     string
	Password  string
	FullName  string
	Role      Role
	Status    Status
	CreatedAt time.Time
}

type Users []User