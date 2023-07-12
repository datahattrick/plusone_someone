package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
)

var DB *database.Queries

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./sql/db/main.db")
	if err != nil {
		log.Fatal(err)
	}

	DB = database.New(db)
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
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
