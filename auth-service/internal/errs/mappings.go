package errs

import "net/http"

func MapErrorCodeToStatus(code ErrorCode) int {
	switch code {
	case CodeUserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
