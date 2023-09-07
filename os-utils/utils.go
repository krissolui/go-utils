package osutils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string, defaultValues ...string) string {
	// check system environment
	env := os.Getenv(key)
	if env != "" {
		return env
	}

	// check .env file
	err := godotenv.Load()
	if err == nil {
		env = os.Getenv(key)
		if env != "" {
			return env
		}
	}

	// return default value if passed
	var defaultValue string
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}

	return defaultValue
}
