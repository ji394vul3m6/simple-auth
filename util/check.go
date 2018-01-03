package util

import "strings"

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
