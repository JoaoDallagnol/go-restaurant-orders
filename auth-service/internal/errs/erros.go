package errs

import (
	"fmt"
)

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

func NewUserNotFound(userID uint) *ResponseError {
	return &ResponseError{
		Code:    CodeUserNotFound,
		Details: fmt.Sprintf("user with ID %d not found", userID),
	}
}

func NewInternalError(details string) *ResponseError {
	return &ResponseError{
		Code:    CodeInternalError,
		Details: details,
	}
}
