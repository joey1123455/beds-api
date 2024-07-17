-- name: CreateLogin2faToken :one
INSERT INTO login2faToken (user_id, token, created_at, expires_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetLogin2faToken :one
SELECT * FROM login2faToken WHERE token = $1 LIMIT 1;

-- name: DeleteLogin2faToken :exec
DELETE FROM login2faToken WHERE token = $1;