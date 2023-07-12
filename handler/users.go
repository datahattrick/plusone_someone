package handler

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
	"github.com/datahattrick/plusone_someone/models"
	"github.com/datahattrick/plusone_someone/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleCreateUser(c *fiber.Ctx) error {
	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
	}

	params := new(parameters)

	if err := c.BodyParser(params); err != nil {
		SendErrorMessage(c, fiber.StatusBadRequest, "Unable to decode json", err)
		return nil
	}

	user, err := utils.Database.DB.CreateUser(c.Context(), database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Username:  params.Username,
		Email:     params.Email,
	})

	if err != nil {
		SendErrorMessage(c, fiber.StatusBadRequest, "Couldn't create user", err)
		return nil
	}

	c.SendStatus(fiber.StatusOK)
	c.JSON(models.DatabaseUserToUser(user))
	return nil
}
