package utils

import (
	"encoding/json"
	"net/http"
)

func MarshalJSONResponse(writer http.ResponseWriter, data any) {
	output, err := json.Marshal(data)
	if err != nil {
		ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(writer, output)
}

func jsonResponse(writer http.ResponseWriter, data []byte) {
	writer.Header().Set("content-type", "application/json")
	writer.Write(data)
}
