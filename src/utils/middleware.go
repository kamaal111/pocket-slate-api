package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-version"
)

type authenticateAppsHeaders struct {
	AppName    string `header:"app-name" bind:"required"`
	AppVersion string `header:"app-version" bind:"required"`
	ApiKey     string `header:"api-key" bind:"required"`
}

func AuthenticateApps(apps []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var headers authenticateAppsHeaders
		err := context.ShouldBindHeader(&headers)
		if err != nil {
			handled := HandleValidationErrors(context, err, "headers")
			if handled {
				return
			}

			ErrorHandler(context, Error{
				Status:  http.StatusBadRequest,
				Message: "Invalid headers provided",
			})
			return
		}

		if headers.ApiKey == "" {
			ErrorHandler(context, Error{
				Status:  http.StatusBadRequest,
				Message: "api-key is required in the headers",
			})
			return
		}

		if headers.AppName == "" {
			ErrorHandler(context, Error{
				Status:  http.StatusBadRequest,
				Message: "app-name is required in the headers",
			})
			return
		}

		if headers.AppVersion == "" {
			ErrorHandler(context, Error{
				Status:  http.StatusBadRequest,
				Message: "app-version is required in the headers",
			})
			return
		}

		headerVersion, err := version.NewVersion(headers.AppVersion)
		if err != nil || headerVersion == nil {
			ErrorHandler(context, Error{
				Status:  http.StatusBadRequest,
				Message: "Invalid version provided",
			})
			return
		}

		rawAppAPIKeys, err := UnwrapEnvironment("APP_API_KEYS")
		if err != nil {
			log.Println("Failed to load app api keys from the environment", err)
			ErrorHandler(context, Error{
				Status:  http.StatusInternalServerError,
				Message: "Something went wrong",
			})
			return
		}

		var appAPIKeys map[string]map[string]string
		err = json.Unmarshal([]byte(rawAppAPIKeys), &appAPIKeys)
		if err != nil {
			log.Println("Failed to unmarshal app api keys", err)
			ErrorHandler(context, Error{
				Status:  http.StatusInternalServerError,
				Message: "Something went wrong",
			})
			return
		}

		tokens := appAPIKeys[headers.AppName]
		var tokenVersions []*version.Version
		for tokenVersion := range tokens {
			parsedVersion, err := version.NewVersion(tokenVersion)
			if err != nil {
				log.Println("Failed to unmarshal app api keys", err)
				ErrorHandler(context, Error{
					Status:  http.StatusInternalServerError,
					Message: "Something went wrong",
				})
				return
			}

			if parsedVersion.GreaterThanOrEqual(headerVersion) {
				tokenVersions = append(tokenVersions, parsedVersion)
			}
		}

		if len(tokenVersions) == 0 {
			ErrorHandler(context, Error{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		sort.Slice(tokenVersions, func(i, j int) bool {
			return tokenVersions[i].LessThan(tokenVersions[j])
		})
		token := tokens[tokenVersions[0].String()]
		if token == "" || token != headers.ApiKey {
			ErrorHandler(context, Error{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
			})
			return
		}

		context.Next()
	}
}
