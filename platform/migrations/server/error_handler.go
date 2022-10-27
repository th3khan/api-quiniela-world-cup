package server

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseServer struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

var ErrorHandler = func(c *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	var message ResponseServer
	message.Message = err.Error()

	if code == fiber.StatusOK {
		message.Error = false
	} else {
		message.Error = true
	}

	return c.Status(code).JSON(message)
}
