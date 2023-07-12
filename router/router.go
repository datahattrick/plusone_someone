package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRouter(app *fiber.App, hostname string, portListen string) {
	// Backend API
	// Version the API
	api := app.Group("/api")
	api.Use(cors.New(cors.Config{
		AllowOrigins: "http://" + hostname + ":" + portListen,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	v1 := api.Group("/v1")
	// User API
	v1.Get("/user", func(c *fiber.Ctx) error { return c.SendString("Welcome User to v1 routing") })
	v1.Get("/user/:userid", func(c *fiber.Ctx) error { return c.SendString("Searching for user unsuccessfully") })

	v1.Post("/user", func(c *fiber.Ctx) error { return c.SendString("Trying really hard to create a user") })
	v1.Post("/user/:userid", func(c *fiber.Ctx) error { return c.SendString("Updating user, one moment") })

	v1.Delete("/user/:userid", func(c *fiber.Ctx) error { return c.SendString("Deleting user, you sure?") })

	//Posts API
	v1.Get("/posts", func(c *fiber.Ctx) error { return c.SendString("Getting all of the posts") })
	v1.Get("/posts/:postid", func(c *fiber.Ctx) error { return c.SendString("Getting one of the posts") })

	v1.Post("/posts", func(c *fiber.Ctx) error { return c.SendString("Creating a post") })

	v1.Delete("/posts/:postid", func(c *fiber.Ctx) error { return c.SendString("Deleting a post") })

	//Other endpoint hits on the api
	api.All("*", func(c *fiber.Ctx) error { return c.SendStatus(404) })

	// Serve the web application
	app.Static("/", "./web/public")
	// Prepare a fallback route to always serve 'index.html'.
	app.Static("*", "./web/public/404.html")
}
