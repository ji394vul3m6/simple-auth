package util

import (
	"regexp"
	"strings"
)

const (
	// UUIDPattern is regex of UUID
	UUIDPattern = "^[0-9a-f]{8}(-[0-9a-f]{4}){4}[0-9a-f]{8}$"
	// MD5Pattern is regex for md5
	MD5Pattern = "^[a-f0-9]{32}$"
)

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

// IsValidMD5 will check if string is standard md5 or not
func IsValidMD5(str string) bool {
	match, _ := regexp.MatchString(MD5Pattern, str)
	return match
}

// IsInSlice will check if check is in container or not
func IsInSlice(check interface{}, container []interface{}) bool {
	for _, v := range container {
		if v == check {
			return true
		}
	}
	return false
}
