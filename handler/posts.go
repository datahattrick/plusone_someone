package handler

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
	"github.com/datahattrick/plusone_someone/models"
	"github.com/datahattrick/plusone_someone/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandlerCreatePost(c *fiber.Ctx) error {
	type parameters struct {
		Message  string `json:"message"`
		AuthorID string `json:"author_id"`
		UserID   string `json:"user_id"`
	}

	params := new(parameters)
	if err := c.BodyParser(params); err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to create post", err)
	}

	posts, err := utils.Database.DB.CreatePost(c.Context(), database.CreatePostParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Message:   params.Message,
		AuthorID:  params.AuthorID,
		UserID:    params.UserID,
	})

	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Couldn't create post", err)
	}

	return c.Status(fiber.StatusOK).JSON(models.DatabasePostToPost(posts))
}

func HandleDeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	err := utils.Database.DB.DeletePost(c.Context(), id)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to delete post: "+id, err)
	}
	return c.SendStatus(fiber.StatusOK)
}

func handleGetPostByUser(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := utils.Database.DB.GetPostsByUser(c.Context(), id)
	if err != nil {
		post, err = utils.Database.DB.GetPostsByAuthor(c.Context(), id)
		if err != nil {
			return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find a post by user: "+id, err)
		}
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabasePostsToPosts(post))
}

func HandleGetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := utils.Database.DB.GetPostByID(c.Context(), id)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find a post by user: "+id, err)
	}

	return c.Status(fiber.StatusOK).JSON(models.DatabasePostToPost(post))
}

func HandleGetPosts(c *fiber.Ctx) error {
	posts, err := utils.Database.DB.GetAllPosts(c.Context())
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find any posts", err)
	}

	return c.Status(fiber.StatusOK).JSON(models.DatabasePostsToPosts(posts))
}
