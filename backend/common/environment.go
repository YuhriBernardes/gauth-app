package common

import (
	"fmt"
	"os"
	"strconv"
)

func EnvIntOrDefault(envName string, defaultValue int) int {
	envValue := os.Getenv(envName)

	if envValue == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(envValue)

	if err != nil {
		panic(fmt.Sprintf("Failed to convert environment %s to integer (%s)", envName, envValue))
	}

	return intValue
}
