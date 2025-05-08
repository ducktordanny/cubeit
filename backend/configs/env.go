package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Production       bool
	Port             string
	RedirectURI      string
	ClientID         string
	ClientSecret     string
	DBName           string
	DBUser           string
	DBPassword       string
	DBHost           string
	DBHostPort       string
	DBDisableSSLMode bool
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Production: getEnvBool("PRODUCTION", false),
		Port:       getEnv("PORT", "6969"),
		RedirectURI: getEnv(
			"REDIRECT_URI",
			"http://localhost:6969/api/v1/oauth/callback",
		),
		ClientID:         getEnv("CLIENT_ID", ""),
		ClientSecret:     getEnv("CLIENT_SECRET", ""),
		DBName:           getEnv("POSTGRES_DB", "cubeit-local"),
		DBUser:           getEnv("POSTGRES_USER", "admin"),
		DBPassword:       getEnv("POSTGRES_PASSWORD", "something"),
		DBHost:           getEnv("POSTGRES_HOST", "localhost"),
		DBHostPort:       getEnv("POSTGRES_HOST_PORT", "5432"),
		DBDisableSSLMode: getEnvBool("POSTGRES_DISABLE_SSL_MODE", false),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return fallback
}
