package models

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
)

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

func DatabaseUsersToUsers(dbusers []database.User) []User {
	users := []User{}
	for _, dbuser := range dbusers {
		users = append(users, DatabaseUserToUser(dbuser))
	}
	return users
}

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Message   string    `json:"message"`
	AuthorID  string    `json:"author_id"`
	UserID    string    `json:"user_id"`
}

func DatabasePostToPost(dbpost database.Post) Post {
	return Post{
		ID:        dbpost.ID,
		CreatedAt: dbpost.CreatedAt,
		UpdatedAt: dbpost.UpdatedAt,
		Message:   dbpost.Message,
		AuthorID:  dbpost.AuthorID,
		UserID:    dbpost.AuthorID,
	}
}

func DatabasePostsToPosts(dbposts []database.Post) []Post {
	posts := []Post{}
	for _, dbpost := range dbposts {
		posts = append(posts, DatabasePostToPost(dbpost))
	}
	return posts
}
