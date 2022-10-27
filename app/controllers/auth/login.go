package auth

import "github.com/gofiber/fiber/v2"

type LoginResponse struct {
	User  string
	Token string
}

func Login(ctx *fiber.Ctx) error {
	return ctx.JSON(LoginResponse{User: "user", Token: "53265323uvdjkvsaduaysd"})
}
