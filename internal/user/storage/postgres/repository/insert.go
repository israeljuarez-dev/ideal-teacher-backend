package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/models"
	sqlc "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/sqlc"
)

func (r *Repository) Insert(ctx context.Context, u *models.InsertUserParams) (*domain.User, error) {
	userParams := sqlc.CreateUserParams{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	ur, err := r.query.CreateUser(ctx, userParams)
	if err != nil {
		return nil, err
	}

	parsedID, err := uuid.FromBytes(ur.ID.Bytes[:])
	if err != nil {
		return nil, err
	}

	var createdAt time.Time
	if ur.CreatedAt.Valid {
		createdAt = ur.CreatedAt.Time
	}

	user := &domain.User{
		ID:        parsedID,
		Email:     ur.Email,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Status:    domain.UserStatus(ur.Status),
		CreatedAt: createdAt,
	}

	return user, nil
}
