package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gorilla/mux"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/routers"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/adapters/storage/postgres"
	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"

)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading configuration", "error", err)
		return
	}

	ctx := context.Background()

	// Conectar a la base de datos
	db, err := postgres.New(ctx, cfg.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", cfg.DB.Name)

	// Migrar tablas
	if err := db.Migrate(cfg.DB); err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully migrated the database")

	router := mux.NewRouter()

	// Configurar tus rutas
	routers.InitRouters(router, db)

	// Iniciar servidor pasando el 'mux' directamente
	if err := config.StartServer(router, cfg.HTTP.Port); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}

	slog.Info("Server initialized successfully", "port", cfg.HTTP.Port)
}
