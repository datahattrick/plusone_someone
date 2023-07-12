package handler

import (
	"net/mail"

	"github.com/datahattrick/plusone_someone/utils"
	"github.com/gofiber/fiber/v2"
)

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Login(c *fiber.Ctx, db *utils.Sqldb) error {

	// type LoginInput struct {
	// 	Identity string `json:"identity"`
	// 	Password string `json:"password"`
	// }

	// input := new(LoginInput)

	// var userData models.User

	// if err := c.BodyParser(&input); err != nil {
	// 	return SendErrorMessage(c, fiber.StatusBadRequest, "Error on login request", err)
	// }

	// identity := input.Identity
	// pass := input.Password

	// if isEmail(identity) {
	// 	user, err = db.DB.GetUserByEmail(c.Context(), identity)
	// }

	// if user == nil {
	// 	return SendErrorMessage(c, fiber.StatusUnauthorized, "user not found", err)
	// }

	return c.SendString("oof")
}
