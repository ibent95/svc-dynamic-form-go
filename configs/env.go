package configs

import (
    "os"
    "log"

    "github.com/joho/godotenv"
)

type EnvConfig struct {
	Driver  string
	URL     string
	AppHost string
	AppPort string
}

var Env EnvConfig

func LoadEnv() {
	// Muat .env ke dalam os.Getenv
	if err := godotenv.Load(); err != nil {
		log.Println("❌ No .env file found or failed to load")
	}

	Env.Driver = getEnv("DATABASE_DRIVER", "")
	Env.URL = getEnv("DATABASE_URL", "")

	if Env.Driver == "" || Env.URL == "" {
		log.Fatal("❌ Missing required environment: DATABASE_DRIVER or DATABASE_URL")
	}

	Env.AppHost = getEnv("APP_URL", "127.0.0.1")
	Env.AppPort = getEnv("APP_PORT", "8080")
}

func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
