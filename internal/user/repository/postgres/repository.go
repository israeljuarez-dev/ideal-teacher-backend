package repository

import (
	"context"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/models"
	sqlc "github.com/israeljuarez-dev/ideal-teacher-backend/internal/user/repository/sqlc"
)

type (
	UserRepository interface {
		Insert(ctx context.Context, in *models.InsertUserParamsIn) (*models.InsertUserOut, error)
		GetByEmail(ctx context.Context, email string) (*models.GetUserByEmailOut, error)
	}

	repository struct {
		db    *postgres.DB
		query *sqlc.Queries
	}
)

func New(db *postgres.DB) UserRepository {
	return &repository{db: db, query: sqlc.New(db.Pool)}
}
