package users

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name" validator:"required"`
	LastName  string    `json:"last_name" validator:"required"`
	Email     string    `json:"email" validator:"required,email"`
	Username  string    `json:"username" validator:"required"`
	ApiKey    string    `json:"api_key"`
}

func DatabaseUserToUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		FirstName: dbuser.FirstName,
		LastName:  dbuser.LastName,
		Email:     dbuser.Email,
		Username:  dbuser.Username,
		ApiKey:    dbuser.ApiKey,
	}
}

func DatabaseUsersToUsers(dbusers []database.User) []User {
	users := []User{}
	for _, dbuser := range dbusers {
		users = append(users, DatabaseUserToUser(dbuser))
	}
	return users
}
