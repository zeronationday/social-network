-- name: ListUsers :many
SELECT * FROM users;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, created_at, updated_at;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;
