package handler

import "github.com/gofiber/fiber/v2"

func SendError(c *fiber.Ctx, code int, msg string) error {
	c.SendStatus(code)
	return c.JSON(fiber.Map{
		"error": msg,
	})
}
