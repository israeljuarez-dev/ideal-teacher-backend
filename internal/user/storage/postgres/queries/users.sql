-- name: CreateUser :one
INSERT INTO users (
    email,
    password,
    first_name,
    last_name
)
VALUES (
    $1, $2, $3, $4
)
RETURNING id, email, first_name, last_name, status, created_at;