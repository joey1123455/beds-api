package common

import "strings"

func IsEmptyOrSpaces(s string) bool {
	return s == "" || len(strings.TrimSpace(s)) == 0
}

func IsValidApplicationStatus(s string) bool {
	val := strings.ToUpper(s)

	switch val {
	case "PENDING", "ACCEPTED", "REJECTED", "ENROLLED", "INCOMPLETE", "APPLICATION_CANCELLED", "APPLICATION_RECEIVED", "AWAITING_PAYMENT", "WAITING_LIST":
		return true
	default:
		return false
	}
}

func IsValidAddmissionStatus(s string) bool {
	val := strings.ToUpper(s)
	switch val {
	case "ACCEPTED", "REJECTED", "PENDING", "DEFERED":
		return true
	}
	return false
}
