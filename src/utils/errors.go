package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type errorMessage struct {
	Message string `json:"message"`
}

func ErrorHandler(context *gin.Context, errorObject Error) {

	context.AbortWithStatusJSON(errorObject.Status, errorMessage{Message: errorObject.Message})
}

func HandleValidationErrors(context *gin.Context, err error, placeOfFailure string) bool {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return false
	}

	for _, err := range validationErrors {
		ErrorHandler(context, Error{
			Status:  http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("'%s' is %s in the %s", PascalToSnakeCase(err.Field()), err.Tag(), placeOfFailure),
		})

		return true
	}

	return false
}
