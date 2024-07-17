-- name: CreateCPasswordResetToken :one
INSERT INTO password_reset_token
    (user_id, token, created_at, expires_at)
VALUES
    ($1, $2, $3, $4)
RETURNING *;

-- name: GetPasswordResetToken :one
SELECT *
FROM password_reset_token
WHERE token = $1
LIMIT 1;

-- name: GetPasswordResetTokenByUserID :one
SELECT *
FROM password_reset_token
WHERE user_id = $1
LIMIT 1;

-- name: DeletePasswordResetToken :exec
DELETE FROM password_reset_token
WHERE token = $1;

