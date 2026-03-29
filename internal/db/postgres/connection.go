package postgres

import (
	"context"
	"fmt"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
	url  string
}

func New(ctx context.Context, cfg *config.DB) (*DB, error) {
	url := buildURL(cfg)

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	// Verifica la conexión a la base de datos
	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{
		db,
		url,
	}, nil
}

func buildURL(cfg *config.DB) string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		cfg.PostgresEnv.User,
		cfg.PostgresEnv.Password,
		cfg.PostgresEnv.Host,
		cfg.PostgresEnv.Port,
		cfg.PostgresEnv.Name,
		cfg.PostgresEnv.SSLMode,
		cfg.PostgresEnv.MinConn,
		cfg.PostgresEnv.MaxConn,
	)
}

// ErrorCode devuelve el código de error específico de PostgreSQL
func (db *DB) ErrorCode(err error) string {
	// Valida de que el error sea del tipo *pgconn.PgError
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

// Close cierra la conexión a la base de datos
func (db *DB) Close() {
	db.pool.Close()
}
