package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/israeljuarez-dev/ideal-teacher-backend/internal/config"
)

var (
	level slog.Level
)

const (
	LevelAll   = slog.Level(-8)
	LevelFatal = slog.Level(12)
)

func New(cfg config.Log) *slog.Logger {
	switch cfg.LogLevel {
	case "all":
		level = LevelAll
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource, // muestra archivo y línea en el log
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				lvl := a.Value.Any().(slog.Level)
				switch {
				case lvl < slog.LevelDebug:
					a.Value = slog.StringValue("ALL")
				case lvl > slog.LevelError:
					a.Value = slog.StringValue("FATAL")
				}
			}
			return a
		},
	}

	var handler slog.Handler
	if cfg.ConsoleDecoration {
		handler = slog.NewTextHandler(os.Stdout, opts) // legible en consola (dev)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts) // JSON estructurado (prod)
	}

	return slog.New(handler)
}

// Fatal loguea y termina el programa
func Fatal(ctx context.Context, logger *slog.Logger, msg string, args ...any) {
	logger.Log(ctx, LevelFatal, msg, args...)
	os.Exit(1)
}