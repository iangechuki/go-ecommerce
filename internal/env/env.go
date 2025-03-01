package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valAInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valAInt
}
func GetBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valAsBool, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return valAsBool
}
