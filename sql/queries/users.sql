-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex')
) -- in sqlc each parameter is interpolated into the function
-- This will create a new function with 4 parameters corresponding to the stuff
RETURNING *; -- Return that record

-- name: GetUserByAPIKey :one
SELECT * FROM users where api_key = $1;