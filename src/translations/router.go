package translations

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) *gin.Engine {
	prefix := "/v1/translations"
	group := engine.Group(prefix)
	group.POST("", makeTranslationHandler)
	group.GET("/supported-locales", getSupportedLocalesHandler)

	return engine
}
