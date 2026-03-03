-- name: GetUserByID :one
SELECT
    id,
    email,
    password,
    full_name,
    role,
    status
FROM users
WHERE id = $1
LIMIT 1;


-- name: GetUserByEmail :one
SELECT
    id,
    email,
    password,
    full_name,
    role,
    status
FROM users
WHERE email = $1
LIMIT 1;


-- name: GetAll :many
SELECT
    id,
    email,
    password,
    full_name,
    role,
    status
FROM users
ORDER BY full_name ASC
LIMIT $1 OFFSET $2;


-- name: CreateUser :one
INSERT INTO users (
    email,
    password,
    full_name,
    role,
    status
)
VALUES ($1, $2, $3, $4, COALESCE($5, 'active'))
RETURNING
    id,
    email,
    password,
    full_name,
    role,
    status;


-- name: UpdateUser :exec
UPDATE users
SET
    email = $2,
    full_name = $3,
    role = $4,
    status = $5
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;