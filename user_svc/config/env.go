package config

import (
	"fmt"
	"os"
)

var (
	PORT      = getEnv("PORT", "5000")
	DB        = getEnv("DB", "user.sqlite")
	TOKEN_EXP = getEnv("TOKEN_EXP", "1h")
	TOKEN_KEY = getEnv("TOKEN_KEY", "secret")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf("Environment variable not found :: %v", name))
}
