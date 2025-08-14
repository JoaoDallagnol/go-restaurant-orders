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

func NewRestaurantNotFound(restaurantId uint) *ResponseError {
	return &ResponseError{
		Code:    CodeRestaurantNotFound,
		Details: fmt.Sprintf("restaurant with Id %d not found", restaurantId),
	}
}

func NewDishNotFound(dishId uint) *ResponseError {
	return &ResponseError{
		Code:    CodeDishNotFound,
		Details: fmt.Sprintf("dish with Id %d not found", dishId),
	}
}

func NewInternalError(details string) *ResponseError {
	return &ResponseError{
		Code:    CodeInternalError,
		Details: details,
	}
}
