package errs

type ErrorCode string

const (
	CodePaymentNotFound              ErrorCode = "PAYMENT_NOT_FOUND"
	CodeInvalidAmount                ErrorCode = "INVALID_AMOUNT"
	CodeInternalError                ErrorCode = "INTERNAL_ERROR"
	CodeOrderServiceIntegrationError ErrorCode = "ORDER_SERVICE_INTEGRATION_ERROR"
)
