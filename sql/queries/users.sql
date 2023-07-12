-- name: CreateUser :one 
INSERT INTO users(id, created_at, updated_at, first_name, last_name, email, username, api_key) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserByFirstname :one
SELECT * FROM users WHERE first_name=?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=?;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username=?;

-- name: GetAllUsers :many
SELECT * FROM users;