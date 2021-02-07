package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Init()  {
	_ = godotenv.Load("../.env")
}

func GetEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}