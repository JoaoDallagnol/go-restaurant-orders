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

func NewOrderNotFound(orderId uint) *ResponseError {
	return &ResponseError{
		Code:    CodeOrderNotFound,
		Details: fmt.Sprintf("order with Id %d not found", orderId),
	}
}

func NewOrderItemNotFound(orderItemId uint) *ResponseError {
	return &ResponseError{
		Code:    CodeOrderItemNotFound,
		Details: fmt.Sprintf("order item with Id %d not found", orderItemId),
	}
}

func NewInternalError(details string) *ResponseError {
	return &ResponseError{
		Code:    CodeInternalError,
		Details: details,
	}
}
