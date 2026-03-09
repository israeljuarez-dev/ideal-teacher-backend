-- name: GetUserByID :one
SELECT
    id,
    email,
    password,
    full_name,
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
    status
FROM users
ORDER BY full_name ASC
LIMIT $1 OFFSET $2;


-- name: CreateUser :one
INSERT INTO users (
    email,
    password,
    full_name,
    status
)
VALUES ($1, $2, $3, COALESCE($4, 'active'))
RETURNING
    id,
    email,
    password,
    full_name,
    status;


-- name: UpdateUser :exec
UPDATE users
SET
    email = $2,
    full_name = $3,
    status = $4
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;


-- name: AddUserRole :exec
INSERT INTO user_roles (user_id, role_id) 
values ($1, $2);


-- name: RemoveUserRole :exec
DELETE FROM user_roles 
WHERE user_id = $1 
AND role_id = $2;