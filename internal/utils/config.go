package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file, loading defaults", err)
		os.Setenv("HOSTNAME", "localhost")
		os.Setenv("PORT", "8000")
	}
}
