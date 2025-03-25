-- name: CreateUser :one
INSERT INTO users (ID, username)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users
WHERE api_key = $1;

-- name: DeleteUserByAPIKey :one
DELETE FROM users
WHERE api_key = $1
RETURNING *;