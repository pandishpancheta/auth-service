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
		TCP_PORT: strings.TrimSuffix(os.Getenv("TCP_PORT"), "\n"),

		DB_HOST: strings.TrimSuffix(os.Getenv("DB_HOST"), "\n"),
		DB_PORT: strings.TrimSuffix(os.Getenv("DB_PORT"), "\n"),
		DB_USER: strings.TrimSuffix(os.Getenv("DB_USER"), "\n"),
		DB_PASS: strings.TrimSuffix(os.Getenv("DB_PASS"), "\n"),
		DB_NAME: strings.TrimSuffix(os.Getenv("DB_NAME"), "\n"),

		JWT_SECRET_KEY:       strings.TrimSuffix(os.Getenv("JWT_SECRET_KEY"), "\n"),
		JWT_EXPIRATION_HOURS: int64(jwtExpiration),
	}
}
