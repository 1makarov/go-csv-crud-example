package config

import (
	"github.com/1makarov/go-csv-crud-example/internal/db"
	"os"
)

type (
	Config struct {
		Environment string
		DB          db.ConfigDB
		HTTP        HTTPConfig
	}

	HTTPConfig struct {
		Host string
		Port string
	}
)

func Init() *Config {
	var cfg Config

	setFromEnv(&cfg)

	return &cfg
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")

	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")

	cfg.Environment = os.Getenv("ENV")
}
