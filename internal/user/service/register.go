package service

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/models"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, in *InsertUserIn) (*InsertUserOut, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	newID := uuid.New()
	pgUUID := pgtype.UUID{
		Bytes: newID,
		Valid: true,
	}

	userParams := &models.InsertUserParamsIn{
		ID:        pgUUID,
		Email:     in.Email,
		Password:  string(hashedPassword),
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	udb, err := s.repo.Insert(ctx, userParams)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	u := &InsertUserOut{
		ID:        udb.ID,
		Email:     udb.Email,
		FirstName: udb.FirstName,
		LastName:  udb.LastName,
		Status:    udb.Status,
		CreatedAt: udb.CreatedAt,
	}

	return u, nil
}
