package config

import (
	"os"

	"github.com/joho/godotenv"
)

func NewConfig() *Config {
	godotenv.Load()

	var DB_HOST string
	if DB_HOST = os.Getenv("DB_HOST"); DB_HOST == "" {
		DB_HOST = "127.0.0.1"
	}

	return &Config{
		Srv: Server{
			Host: "127.0.0.1",
			Port: "5000",
		},
		DB: Database{
			Host:     DB_HOST,
			Port:     "5432",
			Name:     "concrete",
			User:     "postgres",
			Password: "postgres",
		},
	}
}
