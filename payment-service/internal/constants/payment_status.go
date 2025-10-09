package constants

type PaymentStatus string

const (
	APPROVED  PaymentStatus = "APPROVED"
	DECLINED  PaymentStatus = "DECLINED"
	CANCELLED PaymentStatus = "CANCELLED"
)
