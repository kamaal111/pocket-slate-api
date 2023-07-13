package translations

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.Engine {
	prefix := "/v1/translations"
	engine.POST(prefix, makeTranslationHandler)
	engine.GET(filepath.Join(prefix, "supported-locales"), getSupportedLocalesHandler)
	return engine
}
