package postgres

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
)

var (
	//go:embed migrations/*.sql
	migrationsFS embed.FS
)

func (db *DB) Migrate(cfg *config.DB) error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	url := buildMigrateURL(cfg)

	// Crea una nueva instancia de migrate con el driver y la URL de la base de datos
	migrations, err := migrate.NewWithSourceInstance("iofs", driver, url)
	if err != nil {
		return err
	}

	// Aplica las migraciones pendientes
	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func buildMigrateURL(cfg *config.DB) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgresEnv.User,
		cfg.PostgresEnv.Password,
		cfg.PostgresEnv.Host,
		cfg.PostgresEnv.Port,
		cfg.PostgresEnv.Name,
		cfg.PostgresEnv.SSLMode,
	)
}
