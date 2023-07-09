package routers

import (
	"net/http"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	utils.ErrorHandler(w, "Not found", http.StatusNotFound)
}
