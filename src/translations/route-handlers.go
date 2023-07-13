package translations

import (
	"errors"
	"fmt"
	"log"
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

func getSupportedLocalesHandler(context *gin.Context) {
	var query getSupportedLocalesQuery
	err := context.ShouldBindQuery(&query)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, err := range validationErrors {
				context.
					JSON(http.StatusUnprocessableEntity, gin.H{
						"message": fmt.Sprintf("'%s' is %s", utils.PascalToSnakeCase(err.Field()), err.Tag()),
					})
				return
			}
		}

		context.
			JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid query provided",
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
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	if httpErr != nil {
		context.JSON(httpErr.Status, gin.H{
			"message": httpErr.Message,
		})
		return
	}

	context.JSON(http.StatusOK, resp)
}
