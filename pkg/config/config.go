package config

import (
	"os"
	"strconv"
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
		TCP_PORT: os.Getenv("TCP_PORT"),

		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_NAME: os.Getenv("DB_NAME"),

		JWT_SECRET_KEY:       os.Getenv("JWT_SECRET_KEY"),
		JWT_EXPIRATION_HOURS: int64(jwtExpiration),
	}
}
