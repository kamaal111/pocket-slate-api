package health

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.Engine {
	prefix := "/v1/health"
	engine.GET(filepath.Join(prefix, "ping"), pingHandler)
	return engine
}
