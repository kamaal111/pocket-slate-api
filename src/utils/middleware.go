package utils

import (
	"fmt"
	"net/http"
)

func AllowHTTPMethods(methods []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestMethod := r.Method
		for _, method := range methods {
			if requestMethod == method {
				next.ServeHTTP(w, r)
				return
			}
		}

		ErrorHandler(w, fmt.Sprintf("%s not allowed", requestMethod), http.StatusMethodNotAllowed)
	})
}
