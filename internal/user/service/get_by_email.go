package service

import (
	"context"
	"fmt"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/dto"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/mapper"
)

func (s *Service) GetByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	udb, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	userResponse := mapper.ToUserResponse(udb)
	
	return userResponse, nil
}
