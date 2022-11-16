package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/controllers/user/payment"
)

func UserPaymentRoutes(userRouter fiber.Router) {
	r := userRouter.Group("/payments")

	r.Post("/", payment.CreatePayment)
}
