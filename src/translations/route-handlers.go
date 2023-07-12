package translations

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func makeTranslationHandler(context *gin.Context) {
	var payload makeTranslationPayload
	err := context.ShouldBindJSON(&payload)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, err := range validationErrors {
				context.
					JSON(http.StatusUnprocessableEntity, gin.H{
						"message": fmt.Sprintf("%s is %s", utils.PascalToSnakeCase(err.Field()), err.Tag()),
					})
				return
			}
		}

		context.
			JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid payload provided",
			})
		return
	}

	var resp string
	var httpErr *utils.Error
	err = withTranslationService(func(ts translationService) {
		resp, httpErr = ts.Translate(payload.Text, payload.SourceLocale, payload.TargetLocale)
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if httpErr != nil {
		context.JSON(httpErr.Status, gin.H{
			"message": httpErr.Message,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"translated_text": resp})
}

func GetSupportedLocalesHandler(writer http.ResponseWriter, request *http.Request) {
	target, err := utils.UnwrapURLQuery(request, "target")
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var resp []supportedLocale
	var httpErr *utils.Error
	err = withTranslationService(func(ts translationService) {
		resp, httpErr = ts.SupportedLanguages(target)
	})
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if httpErr != nil {
		utils.ErrorHandler(writer, httpErr.Message, httpErr.Status)
		return
	}

	utils.MarshalJSONResponse(writer, resp, http.StatusOK)
}
