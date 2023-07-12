package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	// Attempt to load in a .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file", err)
	}

	return os.Getenv(key)
}
