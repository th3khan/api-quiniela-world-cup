package admin

import "github.com/th3khan/api-quiniela-world-cup/pkg/entities"

type (
	GerUsersResponse struct {
		Users []entities.UserResponse `json:"users"`
		entities.Pagination
	}

	UserBaseRequest struct {
		Name          string `json:"name" validate:"required"`
		Email         string `json:"email" validate:"required,email"`
		Active        bool   `json:"active"`
		RoleId        uint   `json:"role_id" validate:"required"`
		EmailVerified bool   `json:"email_verified"`
		Image         string `json:"image"`
	}

	UserCreateRequest struct {
		UserBaseRequest
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	}

	UserUpdateRequest struct {
		UserBaseRequest
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation" validate:"eqfield=Password"`
	}
)
