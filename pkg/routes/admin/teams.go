package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/admin/teams"
)

func TeamRoutes(adminRoute fiber.Router) {
	teamRoutes := adminRoute.Group("/teams")

	teamRoutes.Get("/", teams.GetTeams)
	teamRoutes.Get("/:id", teams.GetTeam)
	teamRoutes.Post("/", teams.CreateTeam)
	teamRoutes.Put("/:id", teams.UpdateTeam)
	teamRoutes.Delete("/:id", teams.DeleteTeam)
}
