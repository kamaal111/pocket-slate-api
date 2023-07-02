package translations

import "net/http"

func TranslationsRoutes(mux *http.ServeMux, middleware func(next http.Handler) http.Handler) {
	mux.Handle("/v1/translations/supported-locales", middleware(http.HandlerFunc(getSupportedLocalesHandler)))
}
