package service

import (
	"context"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/dto"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/repository"
)

type (
	UserService interface {
		Register(ctx context.Context, u *dto.CreateUserRequest) (*dto.UserResponse, error)
		GetByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
	}

	Service struct {
		repo repository.UserRepository
	}
)

func NewUserService(repo repository.UserRepository) UserService {
	return &Service{repo: repo}
}
