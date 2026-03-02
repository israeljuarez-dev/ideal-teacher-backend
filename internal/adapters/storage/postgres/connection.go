package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
* DB es un contenedor para la conexión a bases de datos PostgreSQL.
* que utiliza pgxpool como controlador de base de datos.
* También contiene una referencia a squirrel.StatementBuilderType.
* que se utiliza para crear consultas SQL compatibles con la sintaxis
* PostgreSQL.
 */
type DB struct {
	Pool         *pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

// New crea una nueva instancia de DB
func New(ctx context.Context, cfg *config.DB) (*DB, error) {
	url := buildURL(cfg)

	log.Println("DB URL:", url)

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	// Verifica la conexión a la base de datos
	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	// Configura el generador de consultas SQL con el formato de marcador de posición adecuado para PostgreSQL
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{
		db,
		&psql,
		url,
	}, nil
}

// buildURL construye la URL de conexión PostgreSQL con pool configurado
func buildURL(cfg *config.DB) string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_min_conns=%d pool_max_conns=%d",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.MinConn,
		cfg.MaxConn,
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
	db.Pool.Close()
}
