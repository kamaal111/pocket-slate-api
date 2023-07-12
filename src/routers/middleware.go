package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func jsonMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Next()
	}
}

func authenticateApps(apps []string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			appName := r.Header.Get("app-name")
			appVersion := r.Header.Get("app-version")
			apiKey := r.Header.Get("api-key")
			if appName == "" || appVersion == "" || apiKey == "" {
				utils.ErrorHandler(w, "Forbidden", http.StatusForbidden)
				return
			}

			rawAppAPIKeys, err := utils.UnwrapEnvironment("APP_API_KEYS")
			if err != nil {
				utils.ErrorHandler(w, "Something went wrong", http.StatusInternalServerError)
				return
			}

			var appAPIKeys map[string]map[string]string
			err = json.Unmarshal([]byte(rawAppAPIKeys), &appAPIKeys)
			if err != nil {
				utils.ErrorHandler(w, "Something went wrong", http.StatusInternalServerError)
				return
			}

			token := appAPIKeys[appName][appVersion]
			if token == "" || token != apiKey {
				utils.ErrorHandler(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
