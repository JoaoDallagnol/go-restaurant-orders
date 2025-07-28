package errs

type ErrorCode string

const (
	CodeUserNotFound  ErrorCode = "UserNotFound"
	CodeInternalError ErrorCode = "InternalError"
)
