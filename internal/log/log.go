package log

import (
	"log/slog"
	"os"
	"strings"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
)

func New(cfg config.Container) *slog.Logger {
	return slog.New(buildHandler(cfg.Log)).With(
		slog.String("app", cfg.App.Name),
	)
}

func buildHandler(cfg config.Log) slog.Handler {
	opts := &slog.HandlerOptions{
		Level:     parseLevel(cfg.LogLevel),
		AddSource: cfg.AddSource,
	}

	if cfg.ConsoleDecoration {
		return slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.NewJSONHandler(os.Stdout, opts)
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
