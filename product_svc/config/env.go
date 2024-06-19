package config

import (
	"fmt"
	"os"
)

var (
	PORT             = getEnv("PORT", "3002")
	DB               = getEnv("DB", "product_svc.sqlite")
	MERCHANT_SVC_URL = getEnv("MERCHANT_SVC_URL", "http://localhost:3001")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf("Environment variable not found :: %v", key))
}
