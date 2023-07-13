package translations

import (
	"github.com/gin-gonic/gin"
	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func Router(engine *gin.Engine) *gin.Engine {
	prefix := "/v1/translations"
	group := engine.Group(prefix)

	group.Use(utils.AuthenticateApps([]string{"pocket-slate"}))

	group.POST("", makeTranslationHandler)
	group.GET("/supported-locales", getSupportedLocalesHandler)

	return engine
}
