package constants

type PaymentStatus string

const (
	StatusAppoved   PaymentStatus = "APPROVED"
	StatusDeclined  PaymentStatus = "DECLINED"
	StatusCancelled PaymentStatus = "CANCELLED"
)
