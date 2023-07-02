package translations

import (
	"encoding/json"
	"net/http"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func getSupportedLocalesHandler(writer http.ResponseWriter, request *http.Request) {
	output, err := json.Marshal(struct {
		Message string `json:"message"`
	}{Message: "pong"})
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("content-type", "application/json")
	writer.Write(output)
}
