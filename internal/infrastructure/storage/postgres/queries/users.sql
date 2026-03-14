-- name: GetUserByID :one
SELECT 
    u.id, 
    u.email, 
    u.full_name, 
    u.status,
    r.name as role_name
FROM users u
INNER JOIN user_roles ur ON u.id = ur.user_id
INNER JOIN roles r ON ur.role_id = r.id
WHERE u.id = $1;


-- name: GetUserByEmail :one
SELECT 
    u.id, 
    u.email,
    u.full_name,
    u.status,
    r.name as role_name
FROM users u
INNER JOIN user_roles ur ON u.id = ur.user_id
INNER JOIN roles r ON ur.role_id = r.id
WHERE u.email = $1 
LIMIT 1;

-- name: GetAll :many
SELECT
    u.id,
    u.email,
    u.full_name,
    u.status,
    r.name AS role_name
FROM users u
INNER JOIN user_roles ur ON u.id = ur.user_id
INNER JOIN roles r ON ur.role_id = r.id
ORDER BY u.full_name ASC
LIMIT $1 OFFSET $2;


-- name: CreateUser :one
INSERT INTO users (email, password, full_name, status)
VALUES ($1, $2, $3, COALESCE($4, 'active'))
RETURNING 
    id, 
    email, 
    full_name, 
    status,
    (SELECT r.name 
     FROM roles r 
     JOIN user_roles ur ON ur.role_id = r.id 
     WHERE ur.user_id = users.id) as role_name;


-- name: UpdateUser :one
UPDATE users u
SET 
    email = $2, 
    full_name = $3, 
    status = $4
WHERE u.id = $1
RETURNING 
    u.id, 
    u.email, 
    u.full_name, 
    u.status,
    (SELECT r.name 
     FROM roles r 
     JOIN user_roles ur ON ur.role_id = r.id 
     WHERE ur.user_id = u.id) as role_name;


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