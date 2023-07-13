package routers

import (
	"github.com/gin-gonic/gin"
)

func jsonMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Next()
	}
}
