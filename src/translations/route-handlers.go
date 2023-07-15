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

//	@Summary	Gets supported locales.
//	@Schemes
//	@Description	Gets all the supported locales that can be used to translate.
//	@Tags			translations
//	@Accept			json
//	@Produce		json
//	@Param			target		query		string	false	"The target language to translate the locales to."	example(it)	default(en)
//
//	@Param			App-Version	header		string	true	"The version of the app."							example(4.2.0)
//	@Param			App-Name	header		string	true	"The name of the app."								example(sup-app)
//	@Param			Api-Key		header		string	true	"API key registered to the app."					example(1234lmao)
//
//	@Success		200			{object}	[]supportedLocaleResponse
//
//	@Failure		400			{object}	utils.errorMessage
//	@Failure		403			{object}	utils.errorMessage
//	@Failure		422			{object}	utils.errorMessage
//	@Failure		500			{object}	utils.errorMessage
//
//	@Router			/translations/supported-locales [get]
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

	var resp []supportedLocaleResponse
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
