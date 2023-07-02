package utils

import (
	"log"
	"os"
)

func UnwrapEnvironment(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s not defined in the environment\n", key)
	}

	return value
}
