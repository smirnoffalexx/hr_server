package config

import (
	"os"
)

type Config struct {
	Environment string

	AuthToken string

	TgBot struct {
		Token string
		URL   string
	}

	Postgres struct {
		HOST, PORT, USER, PASSWORD, DB, SSLMODE string
	}

	Http struct {
		Port string
	}

	Logger struct {
		LOGLVL string
	}
}

func LoadConfig() (*Config, error) {
	cfg := Config{}

	cfg.Postgres.HOST = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.PORT = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.USER = os.Getenv("POSTGRES_USER")
	cfg.Postgres.PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.DB = os.Getenv("POSTGRES_DB")
	cfg.Postgres.SSLMODE = os.Getenv("POSTGRES_SSLMODE")

	cfg.Logger.LOGLVL = os.Getenv("LOGL")

	cfg.Http.Port = os.Getenv("HTTP_PORT")

	cfg.Environment = os.Getenv("ENVIRONMENT")

	cfg.AuthToken = os.Getenv("AUTH_TOKEN")

	cfg.TgBot.Token = os.Getenv("TG_BOT_TOKEN")
	cfg.TgBot.URL = os.Getenv("TG_BOT_URL")

	return &cfg, nil
}
