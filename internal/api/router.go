package api

import (
	"github.com/datahattrick/plusone_someone/internal/posts"
	"github.com/datahattrick/plusone_someone/internal/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Api struct {
	api *fiber.Router
}

// Manage api endpoints

func Swagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  true,
		DocExpansion: "list",
	}))
}

func V1(app *fiber.App) {
	v1 := app.Group("/v1")

	users := v1.Group("/users")
	usersRouter(users)

	posts := v1.Group("/posts")
	postsRouter(posts)
}

func usersRouter(app fiber.Router) {
	app.Get("/", users.GetAllUsers)
	app.Get("/:id", users.GetUserByID)
	app.Post("/", users.CreateUser)
	app.Delete("/:id", users.DeleteUser)
	app.Get("/posts/:id", users.GetPostByUser)
}

func postsRouter(app fiber.Router) {
	app.Get("/", posts.GetAllPosts)
	app.Get("/:id", posts.GetPostByID)
	app.Post("/", posts.CreatePost)
	app.Delete("/:id", posts.DeletePost)
}
