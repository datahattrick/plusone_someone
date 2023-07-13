// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: posts.sql

package database

import (
	"context"
	"time"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, message, author_id, user_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, created_at, updated_at, message, author_id, user_id
`

type CreatePostParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Message   string
	AuthorID  string
	UserID    string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Message,
		arg.AuthorID,
		arg.UserID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Message,
		&i.AuthorID,
		&i.UserID,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE id=?
`

func (q *Queries) DeletePost(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT id, created_at, updated_at, message, author_id, user_id FROM posts
`

func (q *Queries) GetAllPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Message,
			&i.AuthorID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, created_at, updated_at, message, author_id, user_id FROM posts WHERE id=?
`

func (q *Queries) GetPostByID(ctx context.Context, id string) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Message,
		&i.AuthorID,
		&i.UserID,
	)
	return i, err
}

const getPostsByAuthor = `-- name: GetPostsByAuthor :many
SELECT id, created_at, updated_at, message, author_id, user_id FROM posts WHERE author_id=?
`

func (q *Queries) GetPostsByAuthor(ctx context.Context, authorID string) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByAuthor, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Message,
			&i.AuthorID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByUser = `-- name: GetPostsByUser :many
SELECT id, created_at, updated_at, message, author_id, user_id FROM posts WHERE user_id=?
`

func (q *Queries) GetPostsByUser(ctx context.Context, userID string) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Message,
			&i.AuthorID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}