package http

import (
	"github.com/datahattrick/plusone_someone/internal/api"
	"github.com/datahattrick/plusone_someone/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServer(cfg *utils.Config) {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://" + cfg.Server.Host + ":" + cfg.Server.Port + ",http://localhost:3000,http://localhost:8000,http://127.0.0.1:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api.Swagger(app)
	api.V1(app)

	// Serve the web application
	app.Static("/", "./web/build")
	// Prepare a fallback route to always serve 'index.html'.
	app.Static("*", "./tmp/404.html")

	app.Listen(cfg.Server.Host + ":" + cfg.Server.Port)
}
