package http

import "net/http"

type errorResponse struct {
	Message string `json:"message"`
}

func createErrorResponse(status int) *errorResponse {
	var msg string
	switch status {
	case http.StatusBadRequest:
		msg = "Bad Request"
	case http.StatusInternalServerError:
		msg = "Internal Server Error"
	default:
		msg = "Unknown Error"
	}
	return &errorResponse{Message: msg}
}
