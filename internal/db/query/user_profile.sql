-- name: CreateUserProfile :one
INSERT INTO user_profile (
    first_name, last_name, phone_number, country,
    user_id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7
) RETURNING *;

-- name: GetUserProfile :one
SELECT * FROM user_profile WHERE id = $1 LIMIT 1; 

-- name: GetUserProfileByUserID :one
SELECT * FROM user_profile WHERE user_id = $1 LIMIT 1; 