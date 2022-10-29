package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/auth"
)

func AuthRoutes(app fiber.Router) {
	authGroup := app.Group("/auth")

	authGroup.Post("/login", auth.Login)
}
