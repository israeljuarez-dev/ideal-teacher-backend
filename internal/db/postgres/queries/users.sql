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
SELECT 
    u.id, 
    u.email, 
    u.password, 
    u.first_name, 
    u.last_name, 
    u.status, 
    u.created_at,
    r.name AS role_name
FROM users u
LEFT JOIN user_roles ur ON u.id = ur.user_id
LEFT JOIN roles r ON ur.role_id = r.id
WHERE u.email = $1
LIMIT 1;