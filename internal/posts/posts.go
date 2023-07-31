package posts

import (
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
)

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
		UserID:    dbpost.UserID,
	}
}

func DatabasePostsToPosts(dbposts []database.Post) []Post {
	posts := []Post{}
	for _, dbpost := range dbposts {
		posts = append(posts, DatabasePostToPost(dbpost))
	}
	return posts
}
