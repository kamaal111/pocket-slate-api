package translations

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func makeTranslationHandler(context *gin.Context) {
	var payload makeTranslationPayload
	err := context.ShouldBindJSON(&payload)
	if err != nil {
		handled := utils.HandleValidationErrors(context, err, "body")
		if handled {
			return
		}

		utils.ErrorHandler(context, utils.Error{
			Status:  http.StatusBadRequest,
			Message: "Invalid body provided",
		})
		return
	}

	var resp string
	var httpErr *utils.Error
	err = withTranslationService(func(ts translationService) {
		resp, httpErr = ts.Translate(payload.Text, payload.SourceLocale, payload.TargetLocale)
	})
	if err != nil {
		log.Println("Failed to get translation context", err)
		utils.ErrorHandler(context, utils.Error{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	if httpErr != nil {
		utils.ErrorHandler(context, *httpErr)
		return
	}

	context.JSON(http.StatusOK, gin.H{"translated_text": resp})
}

func getSupportedLocalesHandler(context *gin.Context) {
	var query getSupportedLocalesQuery
	err := context.ShouldBindQuery(&query)
	if err != nil {
		handled := utils.HandleValidationErrors(context, err, "query params")
		if handled {
			return
		}

		utils.ErrorHandler(context, utils.Error{
			Status:  http.StatusBadRequest,
			Message: "Invalid query provided",
		})
		return
	}

	var resp []supportedLocale
	var httpErr *utils.Error
	err = withTranslationService(func(ts translationService) {
		resp, httpErr = ts.SupportedLanguages(query.Target)
	})

	if err != nil {
		log.Println("Failed to get supported languages", err)
		utils.ErrorHandler(context, utils.Error{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		return
	}

	if httpErr != nil {
		utils.ErrorHandler(context, *httpErr)
		return
	}

	context.JSON(http.StatusOK, resp)
}
