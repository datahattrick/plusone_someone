package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Attempt to load in a .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file", err)
	}

	// Load Environment config options, no need for the app to crash if a simple port is not set,
	// Configure it to use a default if nothing can be found
	portListen := os.Getenv("HTTP_PORT")
	if portListen == "" {
		portListen = "8000"
		log.Println("Unable to load HTTP_PORT variable, using the default port :", portListen)

	}

	app.Static("/", "./web/public")

	// Prepare a fallback route to always serve 'index.html'.
	app.Static("*", "./web/public/index.html")

	log.Fatal(app.Listen(":" + portListen))
}
