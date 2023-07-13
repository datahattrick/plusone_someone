package handler

import (
	"net/mail"
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

func handleGetUserByUsername(c *fiber.Ctx, id string) error {
	user, err := utils.Database.DB.GetUserByUsername(c.Context(), id)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find user: "+id, err)
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabaseUserToUser(user))
}

func handleGetUserByEmail(c *fiber.Ctx, id string) error {
	user, err := utils.Database.DB.GetUserByEmail(c.Context(), id)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find user: "+id, err)
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabaseUserToUser(user))
}

func handleGetUserByID(c *fiber.Ctx, id string) error {
	user, err := utils.Database.DB.GetUserById(c.Context(), id)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to find user: "+id, err)
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabaseUserToUser(user))
}

func HandleGetUser(c *fiber.Ctx) error {
	user := c.Params("userid")
	if id, err := mail.ParseAddress(user); err == nil {
		return handleGetUserByEmail(c, id.Address)
	} else if id, err := uuid.Parse(user); err == nil {
		return handleGetUserByID(c, id.String())
	} else {
		return handleGetUserByUsername(c, user)
	}
}

func HandleGetAllUsers(c *fiber.Ctx) error {
	user, err := utils.Database.DB.GetAllUsers(c.Context())
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to get list of users", err)
	}
	return c.Status(fiber.StatusOK).JSON(models.DatabaseUsersToUsers(user))
}

func HandleDeleteUser(c *fiber.Ctx) error {
	err := utils.Database.DB.DeleteUser(c.Context(), c.Params("userid"))
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, "Unable to delete user: "+c.Params("userid"), err)
	}
	return c.SendStatus(fiber.StatusOK)
}
