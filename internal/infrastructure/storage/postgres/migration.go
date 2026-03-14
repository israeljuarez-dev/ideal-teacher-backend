package postgres

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"

)

/*
* Variable que guarda es un "contenedor" dentro del binario
* que guarda todos los archivos .sql de la carpeta migrations/.
* Así la aplicación puede ejecutar migraciones sin depender de
* archivos externos en el sistema.
 */
//go:embed migrations/*.sql
var migrationsFS embed.FS

// Migrate ejecuta las migraciones de la base de datos
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

// buildMigrateURL construye la URL de conexión PostgreSQL para migraciones
func buildMigrateURL(cfg *config.DB) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
}
