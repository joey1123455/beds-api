package common

import (
	"log"
	"regexp"
)

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Println("Error compiling regex:", err)
		return false
	}

	return regex.MatchString(email)
}
