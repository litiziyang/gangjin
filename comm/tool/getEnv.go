package tool

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		return
	}
}
