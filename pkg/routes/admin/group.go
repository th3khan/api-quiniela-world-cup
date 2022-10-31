package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/admin/groups"
)

func GroupRoutes(adminRoute fiber.Router) {
	groupRoutes := adminRoute.Group("/groups")

	groupRoutes.Get("/", groups.GetGroups)
	groupRoutes.Get("/:id", groups.GetGroup)
	groupRoutes.Post("/", groups.CreateGroup)
	groupRoutes.Put("/:id", groups.UpdateGroup)
	groupRoutes.Delete("/:id", groups.DeleteGroup)
}
