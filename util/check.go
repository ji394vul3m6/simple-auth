package util

import (
	"regexp"
	"strings"
)

const UUIDPattern = "^[0-9a-f]{8}(-[0-9a-f]{4}){4}[0-9a-f]{8}$"

// IsValidString will check if string in string pointer is valid
func IsValidString(str *string) bool {
	if str == nil {
		return false
	}
	if strings.Trim(*str, " ") == "" {
		return false
	}
	return true
}

// IsValidUUID will check if string is standard uuid or not
func IsValidUUID(str string) bool {
	match, _ := regexp.MatchString(UUIDPattern, str)
	return match
}
