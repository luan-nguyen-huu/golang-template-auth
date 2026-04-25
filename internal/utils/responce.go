package utils

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"response": map[string]interface{}{
			"message": message,
			"code":   statusCode,
			"data":   nil,
		},
	})
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"response": map[string]interface{}{
			"message": message,
			"code":   statusCode,
			"data":   data,
		},
	})
}
