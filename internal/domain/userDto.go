package domain

import (
	"github.com/joey1123455/beds-api/internal/common"
	db "github.com/joey1123455/beds-api/internal/db/sqlc"
	"github.com/joey1123455/beds-api/internal/validator"
)

// type CreateUserRequest struct {
// 	Username        string `json:"username"`
// 	Email           string `json:"email"`
// 	Password        string `json:"password"`
// 	PasswordConfirm string `json:"password_confirm"`
// 	FirstName       string `json:"first_name"`
// 	LastName        string `json:"last_name"`
// 	PhoneNumber     string `json:"phone_number"`
// 	Pin             string `json:"pin"`
// 	Street          string `json:"street"`
// 	City            string `json:"city"`
// 	State           string `json:"state"`
// 	Country         string `json:"country"`
// 	PostalCode      string `json:"postal_code"`
// }

// func (b *CreateUserRequest) Validate(v *validator.Validator) bool {
// 	v.Check(common.ValidateEmail(b.Email), "email", "invalid email address")
// 	v.Check(!common.IsEmptyOrSpaces(b.Username), "username", "username must be provided")
// 	v.Check(common.ValidatePassword(b.Password), "password", "password must be of length 8 and contain at least one uppercase letter, one lowercase letter, one number and one special character")
// 	v.Check(b.Password == b.PasswordConfirm, "password", "password and confirm password must match")
// 	v.Check(common.ValidatePin(b.Pin), "pin", "pin must be six digits")
// 	v.Check(!common.IsEmptyOrSpaces(b.FirstName), "first_name", "first name must be provided")
// 	v.Check(!common.IsEmptyOrSpaces(b.LastName), "last_name", "last name must be provided")
// 	v.Check(!common.IsEmptyOrSpaces(b.Country), "country", "country must be provided")
// 	v.Check(!common.IsEmptyOrSpaces(b.State), "state", "state must be provided")
// 	v.Check(!common.IsEmptyOrSpaces(b.City), "city", "city must be provided")
// 	v.Check(!common.IsEmptyOrSpaces(b.Street), "street", "street must be provided")
// 	v.Check(!common.IsEmptyOrSpaces(b.PostalCode), "postal_code", "postal code must be provided")
// 	v.Check(common.IsValidPhoneNumber(b.PhoneNumber), "phone_number", "invalid phone number format, must be - +234 811 899 7116")
// 	return v.Valid()
// }

type UpdateUserRequest struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
}

func (b *UpdateUserRequest) Validate(v *validator.Validator) bool {
	v.Check(b.ID != "", "id", "id must be provided")
	return v.Valid()
}

type VerifyEmailRequest struct {
	Id         string `json:"id"`
	VerifyCode string `json:"verify_code"`
}

func (b *VerifyEmailRequest) Validate(v *validator.Validator) bool {
	v.Check(b.Id != "", "id", "id must be provided")
	v.Check(b.VerifyCode != "", "verify_code", "verify code must be provided")
	return v.Valid()
}

// type RegisterUserRequest struct {
// 	Email           string `json:"email"`
// 	Password        string `json:"password"`
// 	PasswordConfirm string `json:"password_confirm"`
// }

// func (b *RegisterUserRequest) Validate(v *validator.Validator) bool {
// 	v.Check(common.ValidateEmail(b.Email), "email", "invalid email address")
// 	v.Check(common.ValidePassword(b.Password), "password", "password must be of length 8 and contain at least one uppercase letter, one lowercase letter, one number and one special character")
// 	v.Check(b.Password == b.PasswordConfirm, "password", "password and confirm password must match")
// 	return v.Valid()
// }

type VerifyOtp struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

func (b *VerifyOtp) Validate(v *validator.Validator) bool {
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email address")
	v.Check(!common.IsEmptyOrSpaces(b.Code), "code", "verify code must be provided")
	return v.Valid()
}

type UserProfileResponse struct {
	ID                    string `json:"id"`
	Username              string `json:"username"`
	Email                 string `json:"email"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	PhoneNumber           string `json:"phone_number"`
	Street                string `json:"street"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	Country               string `json:"country"`
	PostalCode            string `json:"postal_code"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	Verified              bool   `json:"verified"`
	RegistrationCompleted bool   `json:"registration_completed"`
	MFAVerified           bool   `json:"mfa_enabled"`
}

func UserResponses(profile db.UserProfile, user db.User) interface{} {
	return UserProfileResponse{
		ID:                    user.ID.String(),
		Username:              user.Username.String,
		Email:                 user.Email,
		FirstName:             profile.FirstName.String,
		LastName:              profile.LastName.String,
		PhoneNumber:           profile.PhoneNumber.String,
		Country:               profile.Country.String,
		City:                  profile.City.String,
		Street:                profile.Street.String,
		State:                 profile.State.String,
		PostalCode:            profile.PostalCode.String,
		Verified:              user.EmailVerified,
		RegistrationCompleted: user.RegistrationCompleted.Bool,
		MFAVerified:           user.MfaEnabled,
		CreatedAt:             user.CreatedAt.String(),
		UpdatedAt:             user.UpdatedAt.String(),
	}
}
