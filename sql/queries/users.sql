-- name: CreateUser :one 
INSERT INTO users(id, created_at, updated_at, first_name, last_name, email, username, api_key) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserBySearch :many
SELECT * FROM users WHERE username like ?1 or
    first_name like ?1 OR
    last_name like ?1 OR
    email like ?1
    ;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=?;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username=?;

-- name: GetUserById :one
SELECT * FROM users WHERE id=?;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=?;