// Provides utility functions for environment variables.

package util

import "os"

// EnvOr returns the value of the environment variable named by the key,
// or the fallback if the variable is not present.
func EnvOr(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
