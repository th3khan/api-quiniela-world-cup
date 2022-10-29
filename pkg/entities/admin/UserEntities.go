package admin

import "github.com/th3khan/api-quiniela-world-cup/pkg/entities"

type (
	GerUsersResponse struct {
		Users []entities.UserResponse
		entities.Pagination
	}
)
