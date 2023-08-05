package http

import (
	"log"

	"github.com/datahattrick/plusone_someone/internal/api"
	"github.com/datahattrick/plusone_someone/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ServerHTTP struct {
	app *fiber.App
}

func NewServerHTTP(cfg *utils.Config) *ServerHTTP {
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

	return &ServerHTTP{
		app: app,
	}
}

func (s *ServerHTTP) Start(cfg *utils.Config) {
	log.Printf("starting server on port %s", cfg.Server.Port)
	if err := s.app.Listen(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
		log.Fatalf("failed to start web server on port %s, %s", cfg.Server.Port, err)
	}
}

func (s *ServerHTTP) ShutDown() error {
	log.Println("Shutting down the server gracefully")
	if err := s.app.Shutdown(); err != nil {
		return err
	}
	log.Println("Server shut down gracefully")
	return nil
}
