package utils

import (
	"fmt"
	"os"
)

func UnwrapEnvironment(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s not defined in the environment", key)
	}

	return value, nil
}
