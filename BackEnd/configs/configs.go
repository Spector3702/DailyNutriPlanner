package configs

import (
	"os"
	"github.com/joho/godotenv"
)

func GetEnvConfigs(key string) string{
	// load .env file
    err := godotenv.Load(".env")
    if err != nil {
        panic("Error loading .env file")
    }
    return os.Getenv(key)
}