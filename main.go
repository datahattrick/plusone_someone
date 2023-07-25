package main

//go:generate swagger generate spec

import (
	"log"

	_ "github.com/datahattrick/plusone_someone/docs"
	"github.com/datahattrick/plusone_someone/router"
	"github.com/datahattrick/plusone_someone/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Plusone Someone API
// @version 0.1
// @description A simple API to create a message and give someone a plusone.
//
// @schemes http https
// @host localhost:8000
// @BasePath /api/v1

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

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

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Locals("allowed"))
		log.Println(c.Params("id"))
		log.Println(c.Query("v"))
		log.Println(c.Cookies("session"))

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	//Setup Routes
	router.SetupRouter(app, hostname, portListen)

	log.Fatal(app.Listen(hostname + ":" + portListen))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
}
