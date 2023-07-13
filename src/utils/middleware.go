package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-version"
)

type authenticateAppsHeaders struct {
	AppName    string `header:"App-Name" bind:"required"`
	AppVersion string `header:"App-Version" bind:"required"`
	ApiKey     string `header:"Api-Key" bind:"required"`
}

func AuthenticateApps(apps []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var headers authenticateAppsHeaders
		err := context.ShouldBindHeader(&headers)
		if err != nil {
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				for _, err := range validationErrors {
					context.
						JSON(http.StatusUnprocessableEntity, gin.H{
							"message": fmt.Sprintf("'%s' is %s in the headers", PascalToSnakeCase(err.Field()), err.Tag()),
						})
					context.Abort()
					return
				}
			}

			context.
				JSON(http.StatusBadRequest, gin.H{"message": "Invalid header provided"})
			context.Abort()
			return
		}

		headerVersion, err := version.NewVersion(headers.AppVersion)
		if err != nil || headerVersion == nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid version provided"})
			context.Abort()
			return
		}

		rawAppAPIKeys, err := UnwrapEnvironment("APP_API_KEYS")
		if err != nil {
			log.Println("Failed to load app api keys from the environment", err)
			context.
				JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
			context.Abort()
			return
		}

		var appAPIKeys map[string]map[string]string
		err = json.Unmarshal([]byte(rawAppAPIKeys), &appAPIKeys)
		if err != nil {
			log.Println("Failed to unmarshal app api keys", err)
			context.
				JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
			context.Abort()
			return
		}

		tokens := appAPIKeys[headers.AppName]
		var tokenVersions []*version.Version
		for tokenVersion := range tokens {
			parsedVersion, err := version.NewVersion(tokenVersion)
			if err != nil {
				log.Println("Failed to unmarshal app api keys", err)
				context.
					JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
				context.Abort()
				return
			}

			if parsedVersion.GreaterThanOrEqual(headerVersion) {
				tokenVersions = append(tokenVersions, parsedVersion)
			}
		}

		if len(tokenVersions) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid version provided"})
			context.Abort()
			return
		}

		sort.Slice(tokenVersions, func(i, j int) bool {
			return tokenVersions[i].LessThan(tokenVersions[j])
		})
		token := tokens[tokenVersions[0].String()]
		if token == "" || token != headers.ApiKey {
			context.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			context.Abort()
			return
		}

		context.Next()
	}
}
