package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/th3khan/api-quiniela-world-cup/pkg/utils"
	"github.com/th3khan/api-quiniela-world-cup/platform/migrations"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if !fiber.IsChild() {
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
