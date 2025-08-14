package errs

import "net/http"

func MapErrorCodeToStatus(code ErrorCode) int {
	switch code {
	case CodeRestaurantNotFound:
		return http.StatusNotFound
	case CodeDishNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
