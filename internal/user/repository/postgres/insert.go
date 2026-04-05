package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/models"
	sqlc "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/sqlc"
)

func (r *Repository) Insert(ctx context.Context, in *models.InsertUserParamsIn) (*models.InsertUserOut, error) {
	userParams := sqlc.CreateUserParams{
		ID:        in .ID,
		Email:     in .Email,
		Password:  in .Password,
		FirstName: in .FirstName,
		LastName:  in .LastName,

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

	u := &models.InsertUserOut{
		ID:        parsedID,
		Email:     ur.Email,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Status:    domain.UserStatus(ur.Status),
		CreatedAt: createdAt,
	}

	return u, nil
}
