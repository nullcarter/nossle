package services

import (
	"encoding/json"
	"net/http"
)

type Response struct{}

type StandardResponse struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
	// Meta *MetaInfo `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// type MetaInfo struct {
// 	Total   int `json:"total,omitempty"`
// 	Page    int `json:"page,omitempty"`
// 	PerPage int `json:"per_page,omitempty"`
// }

func JSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (r Response) Success(w http.ResponseWriter, statusCode int, data any) {
	resp := StandardResponse{
		Success: true,
		Data:    data,
		// Meta: meta
	}

	JSON(w, statusCode, resp)
}

func (r Response) Error(w http.ResponseWriter, statusCode int, code string, message string) {
	resp := StandardResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}

	JSON(w, statusCode, resp)
}
