package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/db/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/router"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
		return
	}

	ctx := context.Background()

	// Conectar a la base de datos
	db, err := postgres.New(ctx, &cfg.Container.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", cfg.Container.DB.PostgresEnv.Name)

	// Migrar tablas
	if err := db.Migrate(&cfg.Container.DB); err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully migrated the database")

	// Configurar tus rutas
	router := router.InitRouters(db)

	// Iniciar servidor
	if err := config.StartServer(router, cfg.Container.App.Port); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}

	slog.Info("Server initialized successfully", "port", cfg.Container.App.Port)
}
