package utils

import (
	"os"
)

func GetEnvOrDefault(key string, defaultValues ...string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}

	// TODO::check .env file

	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}

	return defaultValue
}
