-- name: CreateUser :one
INSERT INTO users (ID, username)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE username = $1
RETURNING *;