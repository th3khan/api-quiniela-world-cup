package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/th3khan/api-quiniela-world-cup/pkg/middleware"
	"github.com/th3khan/api-quiniela-world-cup/pkg/routes"
	"github.com/th3khan/api-quiniela-world-cup/pkg/routes/admin"
	"github.com/th3khan/api-quiniela-world-cup/pkg/routes/user"
	"github.com/th3khan/api-quiniela-world-cup/platform/migrations/server"
)

func CreateServer(port int) {
	app := fiber.New(fiber.Config{
		ErrorHandler: server.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	file := middleware.Logger(app)
	defer file.Close()

	app.Static("/", "./public")

	api := app.Group("/api")

	// routes
	routes.AuthRoutes(api)
	api.Use(middleware.AuthorizationRequired())
	api.Use(middleware.IsUserActive)

	// User app routes.
	user.UserPaymentRoutes(api)

	// Super admin rooutes
	adminRoutes := api.Group("/admin")
	adminRoutes.Use(middleware.IsSuperAdmin)
	admin.UserRoutes(adminRoutes)
	admin.GroupRoutes(adminRoutes)
	admin.TeamRoutes(adminRoutes)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
