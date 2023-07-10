package utils

import (
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ErrorHandler(w http.ResponseWriter, message string, code int) {
	errorResponse := Error{
		Message: message,
		Status:  code,
	}
	log.Printf("failure message: %s; code: %d\n", errorResponse.Message, errorResponse.Status)

	MarshalJSONResponse(w, errorResponse, errorResponse.Status)
}
