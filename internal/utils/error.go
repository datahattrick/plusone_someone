package utils

import "github.com/gofiber/fiber/v2"

func SendErrorMessage(c *fiber.Ctx, code int, msg string, err error) error {
	c.SendStatus(code)
	return c.JSON(fiber.Map{
		"error":   "message",
		"message": msg,
		"data":    err,
	})
}
