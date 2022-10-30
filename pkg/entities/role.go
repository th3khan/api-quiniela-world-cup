package entities

import "github.com/th3khan/api-quiniela-world-cup/app/models"

type (
	RoleResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)

func CreateRoleResponse(role *models.Role) RoleResponse {
	return RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}
