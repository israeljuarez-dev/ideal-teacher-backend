package service

import (
	"context"
	"fmt"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/jwt"
)

func (s *service) Login(ctx context.Context, in *LoginIn) (*LoginOut, error) {
	u, err := s.repo.GetByEmail(ctx, in.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid user or password")
	}

	if err := comparePassword(in.Password, u.Password); err != nil {
		return nil, fmt.Errorf("invalid user or password")
	}

	accessToken, err := jwt.GenerateToken(u, s.cfg)
	if err != nil {
		return nil, fmt.Errorf("error generating token")
	}

	response := &LoginOut{
		Token:     accessToken,
		ExpiresIn: int(s.cfg.ExpirationTime),
	}

	return response, nil
}
