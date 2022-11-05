package utils

import "net/http"

type ErrorResponse struct {
	Message string
	StatusCode int
}

func SendErrorResponse(w http.ResponseWriter, res ErrorResponse) {
	type errorResponse struct {
		Message string `json:"error"`
	}

	SendResponse(w, res.StatusCode, errorResponse{ Message: res.Message })
}