package translations

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.Engine {
	prefix := "/v1/translations"
	engine.POST(prefix, makeTranslationHandler)
	return engine
}
