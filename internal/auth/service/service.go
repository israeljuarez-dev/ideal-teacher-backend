package service

import (
	"context"
	"log/slog"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	repository "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/postgres"
)

type (
	AuthService interface {
		Login(ctx context.Context, in *LoginIn) (*LoginOut, error)
	}

	service struct {
		repo repository.UserRepository
		log  *slog.Logger
		cfg  *config.JWT
	}
)

func New(repo repository.UserRepository, log *slog.Logger, cfg *config.JWT) AuthService {
	return &service{
		repo: repo,
		log: log,
		cfg:  cfg,
	}
}
