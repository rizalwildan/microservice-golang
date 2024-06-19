package config

import (
	"fmt"
	"os"
)

var (
	PORT            = getEnv("PORT", "3001")
	DB              = getEnv("DB", "merchant_svc.sqlite")
	USER_SVC_URL    = getEnv("USER_SVC_URL", "http://localhost:3000")
	PRODUCT_SVC_URL = getEnv("PRODUCT_SVC_URL", "http://localhost:3002")
	PRODUCT_SVC_KEY = getEnv("PRODUCT_SVC_KEY", "wS@oY^DfdCy$CnW7ATHy~9bbiPPctjmV")
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
