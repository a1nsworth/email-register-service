package app

import (
	"log/slog"
	"os"

	"email-register-service/internal/config"
	"email-register-service/internal/connectors/databases/postgres"
)

func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}

func Run() {
	logger := setupLogger().With("run")
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("Error loading config", err)
		os.Exit(1)
	}
	postgres, err := postgres.NewPostgres(cfg.Postgres)
}
