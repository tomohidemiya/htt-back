package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	_ = godotenv.Load("../.env")
}

func GetEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
