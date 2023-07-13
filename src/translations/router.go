package translations

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func Router(engine *gin.Engine, basePath string) *gin.Engine {
	group := engine.Group(filepath.Join(basePath, "translations"))

	group.Use(utils.AuthenticateApps([]string{"pocket-slate"}))

	group.POST("", makeTranslationHandler)
	group.GET("/supported-locales", getSupportedLocalesHandler)

	return engine
}
