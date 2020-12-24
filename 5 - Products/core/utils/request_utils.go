package utils

import (
	"fmt"
	"net/http"
	"products/core/models"
)

var requestErrorStatuses = map[string]models.RequestErrorStats{
	"not-found": {
		Code:    http.StatusNotFound,
		Message: "object not found",
	},
	"inactive-user": {
		Code:    http.StatusForbidden,
		Message: "your account is desactivated. please, contact for the support to activate your account.",
	},
	"missing-credentials": {
		Code:    http.StatusBadRequest,
		Message: "username or password are missing",
	},
	"invalid-credentials": {
		Code:    http.StatusUnauthorized,
		Message: "invalid credentials",
	},
	"internal-error": {
		Code:    http.StatusInternalServerError,
		Message: "an internal error was ocurried. please, contact for the support.",
	},
}

// GetRequestErrorStats returns a code and a message that
// corresponds with the request error ocurried.
func GetRequestErrorStats(err string, aditionalMsg string) models.RequestErrorStats {
	stats := requestErrorStatuses[err]

	if stats.Message == "" {
		stats.Code = http.StatusBadRequest
		stats.Message = fmt.Sprintf("%v.", err)
	}

	if aditionalMsg != "" {
		stats.Message = fmt.Sprintf("%v: %v.", stats.Message, aditionalMsg)
	}

	return stats
}
