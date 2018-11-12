package common

import (
	"os"
)

// GetEnv returns the value of an environment variable or a default value if it
// doesn't exists.
func GetEnv(name string, defValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defValue
	}
	return value
}
