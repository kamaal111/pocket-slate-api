package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func NotFound(context *gin.Context) {
	utils.ErrorHandler(context, utils.Error{Status: http.StatusNotFound, Message: "Not found"})
}
