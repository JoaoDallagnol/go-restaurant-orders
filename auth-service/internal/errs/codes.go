package errs

type ErrorCode string

const (
	CodeUserNotFound     ErrorCode = "USER_NOT_FOUND"
	CodeInternalError    ErrorCode = "INTERNAL_ERROR"
	CodeAuthInvalidCreds ErrorCode = "AUTH_INVALID_CREDENTIALS"
)
