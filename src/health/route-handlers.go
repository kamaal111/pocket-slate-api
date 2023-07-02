package health

import (
	"net/http"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	utils.MarshalJSONResponse(writer, struct {
		Message string `json:"message"`
	}{Message: "pong"})
}
