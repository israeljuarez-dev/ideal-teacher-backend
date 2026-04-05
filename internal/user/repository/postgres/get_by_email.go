package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/models"
)

func (r *Repository) GetByEmail(ctx context.Context, email string) (*models.GetUserByEmailOut, error) {
	ur, err := r.query.GetUserByEmail(ctx, email)
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

	u := &models.GetUserByEmailOut{
		ID:        parsedID,
		Email:     ur.Email,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Password:  ur.Password,
		Status:    domain.UserStatus(ur.Status),
		CreatedAt: createdAt,
		RoleName:  ur.RoleName,
	}

	return u, nil
}
