package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MONGO_PASSWORD         = getEnvOrPanic("MONGO_PASSWORD")
	USER_PASSWORD          = getEnvOrPanic("USER_PASSWORD")
	SESSION_ENCRYPTION_KEY = getEnvOrPanic("SESSION_ENCRYPTION_KEY")
)

func getEnvOrPanic(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		panic("Missing env variable " + key)
	}
	return value
}
