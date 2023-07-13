package health

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.Engine {
	prefix := "/v1/health"
	group := engine.Group(prefix)
	group.GET("/ping", pingHandler)

	return engine
}
