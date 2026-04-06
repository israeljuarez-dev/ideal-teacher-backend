package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/router"
	myValidator "github.com/israeljuarez-dev/ideal-teacher-backend/internal/validator"
	"github.com/israeljuarez-dev/ideal-teacher-backend/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logger.New(config.Log{
			LogLevel:          "debug",
			ConsoleDecoration: true,
		}).Error("Error loading configuration", "error", err)
		return
	}

	// Inicializar Logs
	log := logger.New(cfg.Container.Log)

	ctx := context.Background()

	// Conectar a la base de datos
	db, err := postgres.New(ctx, &cfg.Container.DB)
	if err != nil {
		logger.Fatal(ctx, log, "Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	log.Info("Successfully connected to the database", "db", cfg.Container.DB.PostgresEnv.Name)

	// Migrar tablas
	if err := db.Migrate(&cfg.Container.DB); err != nil {
		logger.Fatal(ctx, log, "Error migrating database", "error", err)
		os.Exit(1)
	}
	log.Info("Successfully migrated the database")

	// Crear validator
	validate := myValidator.New()

	// Configurar tus rutas
	router := router.InitRouters(db, validate, log, &cfg.Container.JWT)

	log.Info("Server starting", "port", cfg.Container.App.Port)

	// Iniciar servidor
	if err := config.StartServer(router, cfg.Container.App.Port); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}

	slog.Info("Server initialized successfully", "port", cfg.Container.App.Port)
}
