package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func jsonMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Next()
	}
}

func allowHTTPMethods(methods []string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			validMethod := utils.FindInSlice(methods, r.Method)
			if validMethod == nil {
				utils.ErrorHandler(w, fmt.Sprintf("%s not allowed", r.Method), http.StatusMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
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

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		observer := &responseObserver{ResponseWriter: w}
		next.ServeHTTP(observer, r)
		elapsed := time.Since(start)
		log.Printf("%d %s: %s in %s", observer.status, r.Method, r.URL.Path, elapsed)
	})
}

type responseObserver struct {
	http.ResponseWriter
	status int
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	o.status = code
}
