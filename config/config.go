package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost              string
	Port                    string
	DBuser                  string
	DBPassword              string
	DBAddress               string
	DBName                  string
	JWTExpirationiInSeconds int64
	JWTSecret               string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBuser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST",
			"127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                  getEnv("DB_NAME", "ecom"),
		JWTExpirationiInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret: getEnv("JWT_SECRET","not-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string,fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i , err := strconv.ParseInt(value,10,64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}