package entities

import "github.com/th3khan/api-quiniela-world-cup/app/models"

type (
	PaymentBase struct {
		Date   string  `json:"date" validate:"required"`
		Amount float32 `json:"amount" validate:"required"`
		Image  string  `json:"image" validate:"required"`
	}

	PaymentRequest struct {
		PaymentBase
	}

	PaymentResponse struct {
		ID     uint  `json:"id"`
		Status int32 `json:"status"`
		PaymentBase
	}
)

func CreatePaymentResponse(payment *models.Payment) PaymentResponse {
	return PaymentResponse{
		PaymentBase: PaymentBase{
			Date:   payment.Date.Format("2006-01-02"),
			Amount: payment.Amount,
			Image:  payment.Image,
		},
		Status: payment.Status,
		ID:     payment.ID,
	}
}
