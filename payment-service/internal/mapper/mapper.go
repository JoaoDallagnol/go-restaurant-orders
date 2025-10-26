package mapper

import (
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/constants"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
)

func MapPaymentToPaymentResponse(payment *model.Payment) model.PaymentResponse {
	return model.PaymentResponse{
		ID:        payment.ID,
		OrderID:   payment.OrderID,
		Amount:    payment.Amount,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt.String(),
	}
}

func MapPaymentListToPaymentResponseList(paymentList *[]model.Payment) []model.PaymentResponse {
	paymentResponseList := make([]model.PaymentResponse, len(*paymentList))

	for i, payment := range *paymentList {
		paymentResponseList[i] = MapPaymentToPaymentResponse(&payment)
	}

	return paymentResponseList
}

func MapPaymentRequestToPayment(payment *model.PaymentRequest, status constants.PaymentStatus) model.Payment {
	return model.Payment{
		OrderID: payment.OrderId,
		Amount:  payment.Amount,
		Status:  status,
	}
}
