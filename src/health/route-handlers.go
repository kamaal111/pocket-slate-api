package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
