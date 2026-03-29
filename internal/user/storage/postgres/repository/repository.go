package repository

import (
	"context"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/domain"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/models"
	sqlc "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/storage/postgres/sqlc"
)

type (
	UserRepository interface {
		Insert(ctx context.Context, u *models.InsertUserParams) (*domain.User, error)
		GetByEmail(ctx context.Context, email string) (*domain.User, error)
	}

	Repository struct {
		db    *postgres.DB
		query *sqlc.Queries
	}
)

func NewUserRepository(db *postgres.DB) UserRepository {
	return &Repository{db: db, query: sqlc.New(db.Pool)}
}
