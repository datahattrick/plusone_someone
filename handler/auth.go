package handler

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return SendError(c, 500, "Logging in is for loosers")
}
