// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_profile.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUserProfile = `-- name: CreateUserProfile :one
INSERT INTO user_profile (
    first_name, last_name, phone_number, country,
    user_id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7
) RETURNING id, first_name, last_name, phone_number, street, city, state, country, postal_code, user_id, created_at, updated_at
`

type CreateUserProfileParams struct {
	FirstName   sql.NullString `json:"first_name"`
	LastName    sql.NullString `json:"last_name"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Country     sql.NullString `json:"country"`
	UserID      uuid.UUID      `json:"user_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (UserProfile, error) {
	row := q.db.QueryRowContext(ctx, createUserProfile,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Country,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i UserProfile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Street,
		&i.City,
		&i.State,
		&i.Country,
		&i.PostalCode,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfile = `-- name: GetUserProfile :one
SELECT id, first_name, last_name, phone_number, street, city, state, country, postal_code, user_id, created_at, updated_at FROM user_profile WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserProfile(ctx context.Context, id uuid.UUID) (UserProfile, error) {
	row := q.db.QueryRowContext(ctx, getUserProfile, id)
	var i UserProfile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Street,
		&i.City,
		&i.State,
		&i.Country,
		&i.PostalCode,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfileByUserID = `-- name: GetUserProfileByUserID :one
SELECT id, first_name, last_name, phone_number, street, city, state, country, postal_code, user_id, created_at, updated_at FROM user_profile WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserProfileByUserID(ctx context.Context, userID uuid.UUID) (UserProfile, error) {
	row := q.db.QueryRowContext(ctx, getUserProfileByUserID, userID)
	var i UserProfile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Street,
		&i.City,
		&i.State,
		&i.Country,
		&i.PostalCode,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}