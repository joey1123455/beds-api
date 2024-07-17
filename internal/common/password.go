package common

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword return bcrypt hashed password using the default cost(10)
func HashPassword(password string) (string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("unable to hash password %v", err)
	}

	return string(bs), nil

}

// CheckPassword checks if plainPassword matches hashedPassword
func CheckPassword(plainPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func ValidatePassword(password string) bool {
	// Check for at least one uppercase letter
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)

	// Check for at least one lowercase letter
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)

	// Check for at least one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

	// Check for at least one special character
	hasSpecial := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(password)

	// Check for minimum length of 8 characters
	hasMinLength := len(password) >= 8

	// Validate the password
	if hasUppercase && hasLowercase && hasDigit && hasSpecial && hasMinLength {
		return true
	} else {
		return false
	}
}

func ValidatePin(password string) bool {
	regexPattern := `^\d{6}$`

	// Compile the regular expression
	regexp := regexp.MustCompile(regexPattern)

	// Test if the string matches the pattern
	if regexp.MatchString(password) {
		return true
	} else {
		return false
	}
}
