package service

import (
	"context"
	"fmt"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/dto"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/jwt"
)

func (s *Service) Login(ctx context.Context, r *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetByEmail(ctx, r.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid user or password")
	}

	if err := comparePassword(r.Password, user.Password); err != nil {
		return nil, fmt.Errorf("invalid user or password")
	}

	accessToken, err := jwt.GenerateToken(user, s.cfg)
	if err != nil {
		return nil, fmt.Errorf("error generating token")
	}

	response := &dto.LoginResponse{
		Token:     accessToken,
		ExpiresIn: int(s.cfg.ExpirationTime),
	}

	return response, nil
}
