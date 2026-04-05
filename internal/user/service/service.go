package service

import (
	"context"

	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
)

type (
	UserService interface {
		Register(ctx context.Context, in *InsertUserIn) (*InsertUserOut, error)
		GetByEmail(ctx context.Context, email string) (*GetByEmailUserOut, error)
	}

	Service struct {
		repo repository.UserRepository
	}
)

func New(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}
