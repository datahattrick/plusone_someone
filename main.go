package main

import (
	"log"

	"github.com/datahattrick/plusone_someone/router"
	"github.com/datahattrick/plusone_someone/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load Environment config options, no need for the app to crash if a simple port is not set,
	// Configure it to use a default if nothing can be found
	portListen := "8000"
	httpPortListen := utils.Config("PORT")
	if httpPortListen == "" {
		log.Println("Unable to load PORT variable, using the default http port :", portListen)
	}
	// Configure hostname
	hostname := utils.Config("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}

	// Connect to the database
	utils.ConnectDB()

	//Setup Routes
	router.SetupRouter(app, hostname, portListen)

	log.Fatal(app.Listen(hostname + ":" + portListen))
}
