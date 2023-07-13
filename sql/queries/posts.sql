-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, message, author_id, user_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id=?;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id=?;

-- name: GetAllPosts :many
SELECT * FROM posts;

-- name: GetPostsByUser :many
SELECT * FROM posts WHERE user_id=?;

-- name: GetPostsByAuthor :many
SELECT * FROM posts WHERE author_id=?;