-- name: CreateUser :one
INSERT INTO users (
    id, 
    email, 
    password, 
    first_name, 
    last_name
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, email, first_name, last_name, status, created_at;

-- name: GetUserByEmail :one
SELECT id, email, password, first_name, last_name, status, created_at
FROM users
WHERE email = $1
LIMIT 1;