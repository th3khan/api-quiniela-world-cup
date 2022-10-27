package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/pkg/middleware"
	"github.com/th3khan/api-quiniela-world-cup/pkg/routes"
	"github.com/th3khan/api-quiniela-world-cup/platform/migrations/server"
)

func CreateServer(port int) {
	app := fiber.New(fiber.Config{
		ErrorHandler: server.ErrorHandler,
	})

	file := middleware.Logger(app)
	defer file.Close()

	// middleware

	// routes
	routes.AuthRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
