package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort   string
	Dsn          string // DB_URL
	AppSecret    string
	ResendApiKey string
	FromMail     string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	resendApiKey := os.Getenv("RESEND_API_KEY")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	fromMail := os.Getenv("FROM_MAIL")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{
		ServerPort:   httpPort,
		Dsn:          dsn,
		AppSecret:    appSecret,
		ResendApiKey: resendApiKey,
		FromMail:     fromMail,
	}, nil
}
