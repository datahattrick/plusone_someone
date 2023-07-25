package handler

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
	"github.com/datahattrick/plusone_someone/models"
	"github.com/datahattrick/plusone_someone/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type search struct {
	searchKey   string
	searchparam string
}

func HandleCreateUser(c *fiber.Ctx) error {
	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
	}

	params := new(parameters)

	if err := c.BodyParser(params); err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to create user", err)
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
		return SendErrorMessage(c, fiber.StatusBadRequest, "Couldn't create user", err)
	}

	c.SendStatus(fiber.StatusOK)
	c.JSON(models.DatabaseUserToUser(user))
	return nil
}

func HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := utils.Database.DB.GetUserById(c.Context(), id)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find user: "+id, err)
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabaseUserToUser(user))
}

func HandleGetUser(c *fiber.Ctx) error {

	if s := c.Query("search"); s != "" {
		user, err := utils.Database.DB.GetUserBySearch(c.Context(), "%"+s+"%")
		if err != nil {
			return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find user: "+s, err)
		}
		return c.Status(fiber.StatusOK).JSON(models.DatabaseUsersToUsers(user))

	} else {
		return c.SendString("no search, no user")
	}

}

// GetUsers godoc
// @Summary Lists all users in the database.
// @Description This will show all users that have been stored in the local DB.These users would have been synced on start up of the application. If not some default users would have been generated for testing.
// @Accept json
// @Produce json
// @Success 200 {object} models.User{}
// @Router /users [get]
func HandleGetAllUsers(c *fiber.Ctx) error {
	user, err := utils.Database.DB.GetAllUsers(c.Context())
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to get list of users", err)
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabaseUsersToUsers(user))
}

func HandleDeleteUser(c *fiber.Ctx) error {
	err := utils.Database.DB.DeleteUser(c.Context(), c.Params("id"))
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to delete user: "+c.Params("id"), err)
	}
	return c.SendStatus(fiber.StatusOK)
}
