package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/th3khan/api-quiniela-world-cup/config"
)

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningMethod:  "HS256",
		SigningKey:     []byte(config.JWTSecret),
	})
}

func AuthSuccess(c *fiber.Ctx) error {
	return c.Next()
}

func AuthError(c *fiber.Ctx, e error) error {
	return fiber.NewError(
		fiber.StatusUnauthorized,
		"Unauthorized",
	)
}
