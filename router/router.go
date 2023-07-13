package router

import (
	"github.com/datahattrick/plusone_someone/handler"
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

	// Good practice to version
	v1 := api.Group("/v1")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) error { return c.SendString("Auth is hard") })

	// User API
	user := v1.Group("/user")
	user.Get("/", handler.HandleGetAllUsers)
	user.Get("/:userid", handler.HandleGetUser)
	user.Post("/", handler.HandleCreateUser)
	//user.Patch("/:userid", func(c *fiber.Ctx) error { return c.SendString("Updating user, one moment") })
	user.Delete("/:userid", handler.HandleDeleteUser)
	user.Get("/post")
	user.Get("/post/:id")

	// Posts API
	posts := v1.Group("/posts")
	posts.Get("/", func(c *fiber.Ctx) error { return c.SendString("Getting all of the posts") })
	posts.Get("/:id", func(c *fiber.Ctx) error { return c.SendString("Getting one of the posts") })
	posts.Post("/", handler.HandlerCreatePost)
	posts.Delete("/:id", handler.HandleDeletePost)

	// Other endpoint hits on the api
	api.All("*", func(c *fiber.Ctx) error { return c.SendStatus(404) })

	// Serve the web application
	app.Static("/", "./web/public")
	// Prepare a fallback route to always serve 'index.html'.
	app.Static("*", "./web/public/404.html")
}
