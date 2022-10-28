package entities

import "github.com/th3khan/api-quiniela-world-cup/app/models"

type (
	UserResponse struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		RoleId uint   `json:"role_id"`
	}
)

func CreateUserResponse(userModel *models.User) UserResponse {
	return UserResponse{
		ID:     userModel.ID,
		Name:   userModel.Name,
		Email:  userModel.Email,
		RoleId: userModel.RoleId,
	}
}
