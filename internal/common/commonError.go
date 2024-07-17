package common

import (
	"errors"
	"fmt"
)

var (
	ErrValidationFailed      string = "validation failed"
	ErrPasswordHashingFailed        = errors.New("failed to hash password")
	ErrPinHashingFailed             = errors.New("failed to hash pin")
	ErrUserCreationFailed           = errors.New("failed to create user")
	ErrUserUserEmailTaken           = errors.New("user email already taken")
	ErrInvalidUUID                  = errors.New("invalid uuid")
	ErrInvalidUserUUID              = errors.New("invalid user id")
	ErrInvalidWalletUUID            = errors.New("invalid wallet id")
	ErrInvalidCurrencyUUID          = errors.New("invalid currency id")

	ErrInvalidVerificationCode           = errors.New("incorrect verification code")
	ErrInvalidPasswordResetToken         = errors.New("incorrect password reset token")
	ErrInvalidPinResetToken              = errors.New("incorrect pin reset token")
	ErrUserEmailNotVerified              = errors.New("user email not verified")
	ErrUserEmailAlreadyVerified          = errors.New("email already verified")
	ErrUserRegistationIncomplete         = errors.New("user yet to finish registration process")
	ErrVerificationCodeExpired           = errors.New("verification code expired")
	ErrPasswordResetCodeExpired          = errors.New("password reset code expired")
	ErrPinResetCodeExpired               = errors.New("pin reset code expired")
	ErrCreatingUser                      = errors.New("failed to create user")
	ErrBadRequest                        = errors.New("bad Request")
	ErrNotExist                          = errors.New("doesn't exist")
	ErrFailToProcess                     = errors.New("fail to process request")
	ErrIncorrectCredentials              = errors.New("incorrect credentials")
	ErrTokenExpired                      = errors.New("token expired")
	ErrTokenParsingFailed                = errors.New("invalid token")
	ErrGeneratingJWTTokens               = errors.New("failed to generate JWT tokens")
	ErrGeneratingPassResetCode           = errors.New("failed to generate password reset token")
	ErrInvalidEmail                      = errors.New("invalid email address")
	ErrUserAlreadyVerified               = errors.New("user already verified")
	ErrFailedToGenerateVerificationCode  = errors.New("failed to generate verification code")
	ErrFailedToGeneratePasswordResetCode = errors.New("failed to generate password reset code")
	ErrFailedToGeneratePinResetCode      = errors.New("failed to generate pin reset code")
	ErrFailedToResetPassword             = errors.New("failed to reset password")
	ErrFailedToResetPin                  = errors.New("failed to reset pin")
	ErrFailedToCreateWallet              = errors.New("failed to create wallet")

	ErrInvalidDepositAmount = errors.New("invalid deposit amount")
	ErrFailedToDeposit      = errors.New("deposit failed")
	ErrFailedToWithdraw     = errors.New("withdrawal failed")
	ErrMustBeFIAT           = errors.New("must be a fiat currency")
)

func NewNodeNotFound(nodeType, nodeID string) error {
	return fmt.Errorf("could not find %s: %s", nodeType, nodeID)
}

func NewUserNotFound(userID string) error {
	return fmt.Errorf("user not found: %v", userID)
}

func NewFailedToUpdateUser(userID string) error {
	return fmt.Errorf("failed to update user: %v", userID)
}

func NewFailedToRequestNewVerificationCode(userID string) error {
	return fmt.Errorf("failed to request new verification code: %v", userID)
}
