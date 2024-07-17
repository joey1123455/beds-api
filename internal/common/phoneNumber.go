package common

import "regexp"

func IsValidPhoneNumber(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\+?[1-9][0-9\s-]{3,15}$`)
	return phoneRegex.MatchString(phone)
}
