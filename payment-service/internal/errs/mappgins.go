package errs

import (
	"net/http"
)

func MapErrorCodeToStatus(code ErrorCode) int {
	switch code {
	case CodePaymentNotFound:
		return http.StatusNotFound
	case CodeInvalidAmount:
		return http.StatusBadRequest
	case CodeInternalError:
		return http.StatusInternalServerError
	case CodeOrderServiceIntegrationError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
