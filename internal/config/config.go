package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Redis struct {
	Host string `env:"HOST,required"`
	Port string `env:"PORT,required"`
}

type Postgres struct {
	Host     string `env:"HOST,required"`
	Port     string `env:"PORT,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	DBName   string `env:"DBNAME,required"`
}

func (p Postgres) GetDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s", p.Host, p.Port, p.User,
		p.Password, p.DBName,
	)
}

type Server struct {
	Host string `env:"HOST,required"`
	Port string `env:"PORT,required"`
}

type Config struct {
	Redis    Redis    `envPrefix:"REDIS_"`
	Postgres Postgres `envPrefix:"POSTGRES_"`
	Server   Server   `envPrefix:"SERVER_"`
}

func NewConfig() (cfg Config, err error) {
	logger := slog.With("config")

	var fileName string
	switch os.Getenv("ENV_TYPE") {
	case "dev":
		fileName = ".dev.env"
	default:
		fileName = ".env"
	}

	err = godotenv.Load(fileName)
	if err != nil {
		logger.Error("Error loading .env file", err)
		return
	}
	if err = env.Parse(&cfg); err != nil {
		logger.Error("Error parsing .env file", err)
		return
	}
	return
}
