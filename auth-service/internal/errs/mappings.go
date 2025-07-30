package errs

import "net/http"

func MapErrorCodeToStatus(code ErrorCode) int {
	switch code {
	case CodeUserNotFound:
		return http.StatusNotFound
	case CodeAuthInvalidCreds:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
