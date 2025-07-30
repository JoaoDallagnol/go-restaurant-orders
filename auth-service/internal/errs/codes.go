package errs

type ErrorCode string

const (
	CodeUserNotFound     ErrorCode = "USER_NOT_FOUND"
	CodeInternalError    ErrorCode = "INTERNAL+ERROR"
	CodeAuthInvalidCreds ErrorCode = "AUTH_INVALID_CREDENTIALS"
)
