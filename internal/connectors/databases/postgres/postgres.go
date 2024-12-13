package postgres

import (
	"log/slog"

	"email-register-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	config config.Postgres
	gorm   *gorm.DB
}

func NewPostgres(config config.Postgres) (*Postgres, error) {
	logger := slog.With("postgres")
	open, err := gorm.Open(postgres.Open(config.GetDsn()))
	if err != nil {
		logger.Error("Error connecting to database", err)
		return nil, err
	}

	return &Postgres{
		config: config,
		gorm:   open,
	}, nil
}
