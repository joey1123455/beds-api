package domain

import (
	"github.com/joey1123455/beds-api/internal/common"
	"github.com/joey1123455/beds-api/internal/validator"
)

type RegisterUserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (b *RegisterUserRequest) Validate(v *validator.Validator) bool {
	v.Check(!common.IsEmptyOrSpaces(b.ConfirmPassword), "confirm_password", "confirm_password must be provided")
	v.Check(common.ValidatePassword(b.Password), "password", "password is not valid - must contain at least 8 characters, one uppercase letter, one lowercase letter, one number and one special character")
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
	v.Check(b.Password == b.ConfirmPassword, "password", "password and confirm password must match")
	return v.Valid()
}

type VerifyEmail struct {
	VerifyCode string `json:"verification_code"`
	Email      string `json:"email"`
}

func (b *VerifyEmail) Validate(v *validator.Validator) bool {
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email address")
	v.Check(!common.IsEmptyOrSpaces(b.VerifyCode), "verification_code", "verification code must be provided")
	return v.Valid()
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *SignInRequest) Validate(v *validator.Validator) bool {
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
	v.Check(!common.IsEmptyOrSpaces(b.Password), "password", "password must be provided")
	return v.Valid()
}

type PasswordResetRequest struct {
	Email           string `json:"email"`
	Code            string `json:"reset_token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (b *PasswordResetRequest) Validate(v *validator.Validator) bool {
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
	v.Check(!common.IsEmptyOrSpaces(b.Code), "reset token", "reset token must be provided")
	v.Check(common.ValidatePassword(b.Password), "password", "password must be of length 8 and contain at least one uppercase letter, one lowercase letter, one number and one special character")
	v.Check(b.Password == b.ConfirmPassword, "password", "password and confirm password must match")
	return v.Valid()
}

// type ResendEmailRequest struct {
// 	Email string `json:"email"`
// }

// func (b *ResendEmailRequest) Validate(v *validator.Validator) bool {
// 	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
// 	return v.Valid()
// }

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

type PasswordReset struct {
	Email           string `json:"email"`
	ResetCode       string `json:"reset_token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"confirm_password"`
}

func (b *PasswordReset) Validate(v *validator.Validator) bool {
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
	v.Check(!common.IsEmptyOrSpaces(b.ResetCode), "reset token", "reset token must be provided")
	v.Check(common.ValidatePassword(b.Password), "password", "password must be of length 8 and contain at least one uppercase letter, one lowercase letter, one number and one special character")
	v.Check(b.Password == b.PasswordConfirm, "password", "password and confirm password must match")
	return v.Valid()
}

type VerifyAuthenticationRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (b *VerifyAuthenticationRequest) Validate(v *validator.Validator) bool {
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
	v.Check(!common.IsEmptyOrSpaces(b.Code), "code", "code must be provided")
	return v.Valid()
}

type CreatePinRequest struct {
	Code       string `json:"verify_code"`
	UserID     string `json:"user_id"`
	Pin        string `json:"pin"`
	ConfirmPin string `json:"confirm_pin"`
}

func (b *CreatePinRequest) Validate(v *validator.Validator) bool {
	v.Check(!common.IsEmptyOrSpaces(b.Code), "verify_code", "verify code must be provided")
	v.Check(common.ValidatePin(b.Pin), "pin", "pin is not valid - must contain 6 digits")
	v.Check(!common.IsEmptyOrSpaces(b.UserID), "user_id", "user_id must be provided")
	v.Check(b.Pin == b.ConfirmPin, "pin", "pin and confirm pin must match")
	return v.Valid()
}

type CreateProfileRequest struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Pin         string `json:"pin"`
	PhoneNumber string `json:"phone_number"`
	Country     string `json:"country"`
}

func (b *CreateProfileRequest) Validate(v *validator.Validator) bool {
	v.Check(!common.IsEmptyOrSpaces(b.FirstName), "first_name", "first_name must be provided")
	v.Check(!common.IsEmptyOrSpaces(b.LastName), "last_name", "last_name must be provided")
	v.Check(common.ValidateEmail(b.Email), "email", "invalid email")
	v.Check(!common.IsEmptyOrSpaces(b.Country), "country", "country must be provided")
	v.Check(common.IsValidPhoneNumber(b.PhoneNumber), "phone_number", "invalid phone number format, must be - +234 811 899 7116")
	v.Check(common.ValidatePin(b.Pin), "pin", "pin is not valid - must contain 6 digits")
	return v.Valid()
}
