package util

import (
	"os"
	"strconv"
)

// GetStrEnv will get string value from env with key.
// If env not set, return fallback
func GetStrEnv(key string, fallback string) string {
	return getEnv(key, fallback)
}

// GetIntEnv will get int value from env with key.
// If env not set, return fallback
func GetIntEnv(key string, fallback int) int {
	strVal := getEnv(key, "")
	val, err := strconv.ParseInt(strVal, 10, 32)
	if err != nil {
		return fallback
	}
	return int(val)
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
