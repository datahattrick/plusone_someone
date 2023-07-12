package handler

import (
	"context"
	"log"

	"github.com/datahattrick/plusone_someone/models"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	log.Println(models.DB.GetUserByEmail(context.Background(), "test@test.com"))
	return c.SendString("oof")
}
