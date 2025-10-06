package errs

import "fmt"

type ResponseError struct {
	Code    ErrorCode `json:"error"`
	Details string    `json:"details"`
}

func (e *ResponseError) Error() string {
	return e.Details
}

type CodedError interface {
	error
	GetCode() ErrorCode
	GetDetails() string
}

func (e *ResponseError) GetCode() ErrorCode {
	return e.Code
}

func (e *ResponseError) GetDetails() string {
	return e.Details
}

func NewPaymentNotFound(paymentId uint) *ResponseError {
	return &ResponseError{
		Code:    CodePaymentNotFound,
		Details: fmt.Sprintf("payment with id %d not found", paymentId),
	}
}

func NewInvalidAmount(amount string) *ResponseError {
	return &ResponseError{
		Code:    CodeInvalidAmount,
		Details: fmt.Sprintf("invalid amount of %s", amount),
	}
}

func NewInternalError(details string) *ResponseError {
	return &ResponseError{
		Code:    CodeInternalError,
		Details: details,
	}
}

func NewOrderServiceIntegrationError() *ResponseError {
	return &ResponseError{
		Code:    CodeOrderServiceIntegrationError,
		Details: "Something went wrong when trying to communicate with order-service",
	}
}
