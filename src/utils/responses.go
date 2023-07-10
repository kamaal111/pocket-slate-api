package utils

import (
	"encoding/json"
	"net/http"
)

func MarshalJSONResponse(writer http.ResponseWriter, data any, statusCode int) {
	output, err := json.Marshal(data)
	if err != nil {
		ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(writer, output, statusCode)
}

func jsonResponse(writer http.ResponseWriter, data []byte, statusCode int) {
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(data)
}
