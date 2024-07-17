package common

import (
	"fmt"
	"time"
)

func IsValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return false
	}
	return true
}

func ValidateDate(dateStr string) (time.Time, error) {
	parseDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parseDate, nil
}

func ParseDate(startDateString string) (time.Time, error) {
	formats := []string{
		"2006-01-02T15:04:05.999999999Z07:00", // Full format with nanoseconds and timezone
		"2006-01-02T15:04:05Z07:00",           // Format with seconds and timezone
		"2006-01-02",                          // Date only
	}

	var parsedTime time.Time
	var err error
	for _, layout := range formats {
		parsedTime, err = time.Parse(layout, startDateString)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("start_date format error: %v", err)
}
