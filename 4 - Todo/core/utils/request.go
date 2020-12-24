package utils

import (
	"fmt"
	"net/http"
)

// RequestErrorStats type struct.
type RequestErrorStats struct {
	Code    int
	Message string
}

// GetRequestErrorStats returns code and message that
// corresponds with the request error occurried.
func GetRequestErrorStats(code string, aditionalMsg string) RequestErrorStats {
	data := map[string]RequestErrorStats{
		"no-query": {
			Code:    http.StatusBadRequest,
			Message: "need query in request body",
		},
		"json-parse-error": {
			Code:    http.StatusBadRequest,
			Message: "error parsing JSON request body",
		},
		"not-found": {
			Code:    http.StatusNotFound,
			Message: "not found",
		},
	}

	stats := data[code]

	if stats.Message == "" {
		stats.Code = http.StatusBadRequest
	}

	if aditionalMsg != "" {
		stats.Message = fmt.Sprintf("%v: %v", stats.Message, aditionalMsg)
	}

	return stats
}
