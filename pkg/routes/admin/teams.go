package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/admin/groups"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/admin/teams"
)

func TeamRoutes(adminRoute fiber.Router) {
	teamRoutes := adminRoute.Group("/teams")

	teamRoutes.Get("/", teams.GetTeams)
	teamRoutes.Get("/:id", groups.GetGroup)
	teamRoutes.Post("/", groups.CreateGroup)
	teamRoutes.Put("/:id", groups.UpdateGroup)
	teamRoutes.Delete("/:id", groups.DeleteGroup)
}
