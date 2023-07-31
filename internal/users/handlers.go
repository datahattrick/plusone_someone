package users

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
	"github.com/datahattrick/plusone_someone/internal/posts"
	"github.com/datahattrick/plusone_someone/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userparams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

// @id				CreateUser
// @tags			user
// @Summary		Create a User account
// @Description	Creates a user account returns the account details
// @Accept			json
// @Produce		json
// @Param			request	body		userparams	true	"User parameters"
// @Success		200		{object}	User{}
// @Router			/users [post]
func CreateUser(c *fiber.Ctx) error {
	params := new(userparams)

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

// @id				GetUsers
// @tags			user
// @Summary		Lists all users in the database.
// @Description	This will show all users that have been stored in the local DB.These users would have been synced on start up of the application. If not some default users would have been generated for testing.
// @Accept			json
// @Produce		json
// @Success		200	{object}	User{}
// @Router			/users [get]
func GetAllUsers(c *fiber.Ctx) error {
	user, err := utils.Database.DB.GetAllUsers(c.Context())
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

// @id				GetPostByUser
// @tags			user
// @Summary		Get all posts created by a user
// @Description	Get all the posts created using a users ID
// @Accept			json
// @Produce		json
// @Param			id	path	string	false	"User ID"
// @Success		200
// @Router			/users/post/{id} [get]
func GetPostByUser(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := utils.Database.DB.GetPostsByUser(c.Context(), id)
	if err != nil {
		post, err = utils.Database.DB.GetPostsByAuthor(c.Context(), id)
		if err != nil {
			return utils.SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find a post by user: "+id, err)
		}
	}
	return c.Status(fiber.StatusOK).JSON(posts.DatabasePostsToPosts(post))
}
