package health

import "net/http"

func HealthRoutes(mux *http.ServeMux, middleware func(next http.Handler) http.Handler) {
	mux.Handle("/v1/health/ping", middleware(http.HandlerFunc(pingHandler)))
}
