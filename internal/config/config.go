package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required"`
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
}

// todo server, cache, s3 config

func MustLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки .env: %s", err)
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("ошибка загрузки .env: %s", err)
	}

	return cfg
}
