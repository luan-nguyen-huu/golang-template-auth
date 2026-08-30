package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type APIWrapper struct {
	Response APIResponse `json:"response"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	WriteJSON(w, statusCode, APIWrapper{
		Response: APIResponse{
			Code:    statusCode,
			Message: message,
			Data:    data,
		},
	})
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	WriteJSON(w, statusCode, APIWrapper{
		Response: APIResponse{
			Code:    statusCode,
			Message: message,
			Data:    nil,
		},
	})
}
