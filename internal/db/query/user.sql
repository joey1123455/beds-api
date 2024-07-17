-- name: CreateUser :one
INSERT INTO users (
    username, email, password, 
    verify_code, code_expire_time, pin, 
    user_role, created_at, updated_at)
VALUES (
    $1, $2, $3, $4, 
    $5, $6, $7, $8,
    $9
    ) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE($2, username),
    email = COALESCE($3, email),
    updated_at = COALESCE($4, updated_at),
    user_role = COALESCE($5, user_role)
WHERE id = $1
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: ActivateUser :exec
UPDATE users
SET email_verified = TRUE, verify_code = NULL, code_expire_time = NULL, 
    updated_at = $2
WHERE id = $1;

-- name: UpdateRegistrationStatus :exec
UPDATE users
SET registration_completed = $2, 
    updated_at = $3
WHERE id = $1;

-- name: UserEnableMfa :exec
UPDATE users
SET mfa_enabled = TRUE, updated_at = $2 
WHERE id = $1;

-- name: ChangeUserName :exec
UPDATE users
SET username = COALESCE($2, username), updated_at = COALESCE($3, updated_at)
WHERE id = $1;

-- name: ChangeUserPassword :exec
UPDATE users
SET password = $2, updated_at = $3
WHERE id = $1;

-- name: ChangeUserVerifyCode :one
UPDATE users
SET verify_code = $2, code_expire_time = $3, updated_at = $4
WHERE id = $1
RETURNING verify_code;

-- name: ChangeUserPin :one
UPDATE users
SET pin = $2, updated_at = $3
WHERE id = $1
RETURNING pin;
