package service

import (
	"context"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/auth/dto"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/repository"
)

type (
	AuthService interface {
		Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	}

	Service struct {
		repo repository.UserRepository
		cfg  *config.JWT
	}
)

func NewAuthService(repo repository.UserRepository, cfg *config.JWT) AuthService {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}
