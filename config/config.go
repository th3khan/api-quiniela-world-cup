package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	return os.Getenv(key)
}

var JWTSecret string

func InitJWT() {
	JWTSecret = Config("APP_SECRET")
}
