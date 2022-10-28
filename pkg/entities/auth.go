package entities

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	TokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	LoginResponse struct {
		User UserResponse
		TokenResponse
	}
)
