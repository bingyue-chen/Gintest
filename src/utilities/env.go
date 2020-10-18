package utilities

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func Getenv(key string) string {
	return os.Getenv(key)
}
