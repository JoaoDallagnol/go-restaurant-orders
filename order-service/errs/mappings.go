package errs

import "net/http"

func MapErroCodeToStatus(code ErrorCode) int {
	switch code {
	case CodeOrderNotFound:
		return http.StatusNotFound
	case CodeOrderItemNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
