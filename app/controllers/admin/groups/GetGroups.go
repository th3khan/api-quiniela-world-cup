package groups

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/constants"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func GetGroups(ctx *fiber.Ctx) error {
	pageQuery := ctx.Query("page")
	perPageQuery := ctx.Query("per_page")

	var page int
	var perPage int
	var err error

	if len(pageQuery) > 0 {
		page, err = strconv.Atoi(pageQuery)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Parámetro {page} es inválido")
		}

		if len(perPageQuery) == 0 {
			perPage = constants.PAGINATION_PER_PAGE_DEFAULT
		} else {
			perPage, err = strconv.Atoi(perPageQuery)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Parámetro {per_page} es inválido")
			}
		}
	}

	db := database.Connection()
	repo := repositories.NewGroupRepository(db)

	groups, total := repo.GetGroups(page, perPage)

	return ctx.JSON(entities.CreateGroupResponseWithPagination(
		page,
		perPage,
		total,
		groups,
	))
}
