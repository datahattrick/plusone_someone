package users

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
	"github.com/datahattrick/plusone_someone/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Userparams struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
}

// @id				CreateUser
// @tags			user
// @Summary		Create a User account
// @Description	Creates a user account returns the account details
// @Accept			json
// @Produce		json
// @Param			request	body		Userparams	true	"User parameters"
// @Success		200		{object}	User{}
// @Router			/users [post]
func CreateUser(c *fiber.Ctx) error {
	params := new(Userparams)

	if err := c.BodyParser(params); err != nil {
		return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Unable to create user", err)
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
		return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Couldn't create user", err)
	}

	c.SendStatus(fiber.StatusOK)
	c.JSON(DatabaseUserToUser(user))
	return nil
}

// @id				GetUserByID
// @tags			user
// @Summary		Return a user by ID
// @Description	Return a single user using their ID
// @Accept			json
// @Produce		json
// @Param			id	path		string	false	"User ID"
// @Success		200	{object}	User{}
// @Router			/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := utils.Database.DB.GetUserById(c.Context(), id)
	if err != nil {
		return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find user: "+id, err)
	}
	return c.Status(fiber.StatusOK).JSON(DatabaseUserToUser(user))
}

func GetAllUsers(c *fiber.Ctx) error {
	user, err := utils.Database.DB.GetAllUsers(c.Context())
	if err != nil {
		return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Unable to get list of users", err)
	}
	return c.Status(fiber.StatusOK).JSON(DatabaseUsersToUsers(user))
}

// @id				GetUserBySearch
// @tags			user
// @Summary		Finds users
// @Description	This will conduct a fuzzy search for a user by checking both first_name, last_name, email and username
// @Accept			json
// @Produce		json
// @Param	search query string false "User"
// @Success		200	{object}	[]User{}
// @Router			/users [get]
func GetUserBySearch(c *fiber.Ctx) error {
	search := c.Query("search")
	if search == "" {
		GetAllUsers(c)
	}
	user, err := utils.Database.DB.GetUserBySearch(c.Context(), "%"+search+"%")
	if err != nil {
		return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Unable to get list of users", err)
	}
	return c.Status(fiber.StatusOK).JSON(DatabaseUsersToUsers(user))
}

// @id				DeleteUser
// @tags			user
// @Summary		Deletes A user
// @Description	This will delete a user, more for cleaning database.
// @Accept			json
// @Produce		json
// @Param			id	path		string	false	"User ID"
// @Success		200	{string}	string	"ok"
// @Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	err := utils.Database.DB.DeleteUser(c.Context(), c.Params("id"))
	if err != nil {
		return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Unable to delete user: "+c.Params("id"), err)
	}
	return c.SendStatus(fiber.StatusOK)
}
