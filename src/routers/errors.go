package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
}
