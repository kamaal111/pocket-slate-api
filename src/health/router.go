package health

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine, basePath string) *gin.Engine {
	group := engine.Group(filepath.Join(basePath, "health"))
	group.GET("/ping", pingHandler)

	return engine
}
