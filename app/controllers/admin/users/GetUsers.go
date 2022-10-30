package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/constants"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities/admin"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func GetUsers(ctx *fiber.Ctx) error {
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
	userRepository := repositories.NewUserRespository(db)

	users := userRepository.GetUsers(page, perPage)
	total := userRepository.GetTotalUsers()

	return ctx.JSON(createUsersResponseWithPagination(page, perPage, total, users))
}

func createUsersResponseWithPagination(page int, perPage int, total int, users []models.User) admin.GerUsersResponse {
	var response admin.GerUsersResponse

	response.Page = page
	response.PerPage = perPage
	response.Total = total

	var usersresponse []entities.UserResponse

	for _, user := range users {
		usersresponse = append(usersresponse, entities.CreateUserResponse(&user))
	}

	response.Users = usersresponse

	return response
}
