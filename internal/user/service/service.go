package service

import (
	"context"
	"log/slog"

	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
)

type (
	UserService interface {
		Register(ctx context.Context, in *InsertUserIn) (*InsertUserOut, error)
		GetByEmail(ctx context.Context, email string) (*GetByEmailUserOut, error)
	}

	Service struct {
		repo repository.UserRepository
		log  *slog.Logger
	}
)

func New(repo repository.UserRepository, log *slog.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}
