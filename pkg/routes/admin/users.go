package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/admin/users"
)

func UserRoutes(adminRouter fiber.Router) {
	userGroup := adminRouter.Group("/users")

	userGroup.Get("/", users.GetUsers)
	userGroup.Get("/:id", users.GetUser)
	userGroup.Post("/", users.CreateUser)
	userGroup.Put("/:id", users.UpdateUser)
	userGroup.Delete("/:id", users.DeleteUser)
}
