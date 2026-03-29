package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
)

func (r *Repository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
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

	user := &domain.User{
		ID:        parsedID,
		Email:     ur.Email,
		Password:  ur.Password, // necesario para comparar hash en login
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Status:    domain.UserStatus(ur.Status),
		CreatedAt: createdAt,
	}

	return user, nil
}
