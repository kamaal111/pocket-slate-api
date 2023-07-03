package translations

import (
	"net/http"
	"strings"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func TranslationsRoutes(mux *http.ServeMux, middleware func(next http.Handler) http.Handler) {
	urlPrefix := "/v1/translations"
	mux.Handle(
		urlPrefix,
		middleware(utils.AllowHTTPMethods([]string{"POST"}, http.HandlerFunc(makeTranslationHandler))),
	)
	mux.Handle(
		strings.Join([]string{urlPrefix, "supported-locales"}, "/"),
		middleware(utils.AllowHTTPMethods([]string{"GET"}, http.HandlerFunc(getSupportedLocalesHandler))),
	)
}
