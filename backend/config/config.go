package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	Env                string
	Port               string
	DBPort             string
	DBName             string
	DBUser             string
	DBPassword         string
	JWTSecret          string
	JWTExpirationHours int
}

var AppConfig *Config

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		DBHost:             getEnv("DB_HOST", "localhost"),
		Port:               getEnv("PORT", "8080"),
		Env:                getEnv("APP_ENV", "development"),
		DBPort:             getEnv("DB_PORT", "3306"),
		DBName:             getEnv("DB_NAME", "sociomile"),
		DBUser:             getEnv("DB_USER", "root"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		JWTSecret:          getEnv("JWT_SECRET", "secret"),
		JWTExpirationHours: 24,
	}

	return AppConfig
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
