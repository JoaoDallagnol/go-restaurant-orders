package errs

import "net/http"

func MapErrorCodeToStatus(code ErrorCode) int {
	switch code {
	case CodeOrderNotFound:
		return http.StatusNotFound
	case CodeOrderItemNotFound:
		return http.StatusNotFound
	case CodeMenuServiceIntegrationError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
