-- name: CreateUser :one 
INSERT INTO users(id, created_at, updated_at, first_name, last_name, email, username, api_key) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserByFirstname :one
SELECT * FROM users WHERE first_name=?;