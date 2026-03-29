package service

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/dto"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/models"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/mapper"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, u *dto.CreateUserRequest) (*dto.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	newID := uuid.New()
	pgUUID := pgtype.UUID{
		Bytes: newID,
		Valid: true,
	}

	userParams := &models.InsertUserParams{
		ID:        pgUUID,
		Email:     u.Email,
		Password:  string(hashedPassword),
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	udb, err := s.repo.Insert(ctx, userParams)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	userResponse := mapper.ToUserResponse(udb)

	return userResponse, nil
}
