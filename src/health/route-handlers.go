package health

import (
	"net/http"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func PingHandler(writer http.ResponseWriter, request *http.Request) {
	utils.MarshalJSONResponse(writer, struct {
		Message string `json:"message"`
	}{Message: "pong"}, http.StatusOK)
}
