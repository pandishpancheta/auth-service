package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	TCP_PORT string

	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string

	JWT_SECRET_KEY       string
	JWT_EXPIRATION_HOURS int64
}

func LoadConfig() *Config {
	jwtExpiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		panic(err)
	}

	return &Config{
		TCP_PORT: strings.TrimSpace(os.Getenv("TCP_PORT")),

		DB_HOST: strings.TrimSpace(os.Getenv("DB_HOST")),
		DB_PORT: strings.TrimSpace(os.Getenv("DB_PORT")),
		DB_USER: strings.TrimSpace(os.Getenv("DB_USER")),
		DB_PASS: strings.TrimSpace(os.Getenv("DB_PASS")),
		DB_NAME: strings.TrimSpace(os.Getenv("DB_NAME")),

		JWT_SECRET_KEY:       strings.TrimSpace(os.Getenv("JWT_SECRET_KEY")),
		JWT_EXPIRATION_HOURS: int64(jwtExpiration),
	}
}
