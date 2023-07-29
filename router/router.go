package router

import (
	"github.com/datahattrick/plusone_someone/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func SetupRouter(app *fiber.App, hostname string, portListen string) {
	// Backend API
	// Version the API
	api := app.Group("/api")
	api.Use(cors.New(cors.Config{
		AllowOrigins: "http://" + hostname + ":" + portListen + ",http://localhost:3000,http://localhost:8000,http://127.0.0.1:8000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Good practice to version
	v1 := api.Group("/v1")

	v1.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  true,
		DocExpansion: "list",
	}))

	// User API
	v1.Get("/users", handler.HandleGetAllUsers)
	user := v1.Group("/user")
	user.Get("/", handler.HandleGetUser) //TODO need middleware
	user.Get("/:id", handler.HandleGetUserByID)
	user.Post("/", handler.HandleCreateUser)
	user.Delete("/:id", handler.HandleDeleteUser)
	// user.Get("/post", ) - TODO: need middlewareAuth
	user.Get("/post/:id", handler.HandleGetPostByUser)

	// Posts API
	posts := v1.Group("/posts")
	posts.Get("/", handler.HandleGetPosts)
	posts.Get("/:id", handler.HandleGetPostByID)
	posts.Post("/", handler.HandlerCreatePost)
	posts.Delete("/:id", handler.HandleDeletePost)

	// Other endpoint hits on the api
	api.All("*", func(c *fiber.Ctx) error { return c.SendStatus(404) })

	// Serve the web application
	app.Static("/", "./web/build")
	// Prepare a fallback route to always serve 'index.html'.
	app.Static("*", "./tmp/404.html")
}
