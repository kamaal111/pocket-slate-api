package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kamaal111/pocket-slate-api/src/health"
	"github.com/kamaal111/pocket-slate-api/src/translations"
	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func Start() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		port, err := utils.UnwrapEnvironment("PORT")
		if err != nil {
			log.Fatal(err)
		}
		serverAddress = fmt.Sprintf(":%s", port)
	}

	mux := http.NewServeMux()

	mux.Handle(
		"/v1/health/ping",
		loggerMiddleware(
			allowHTTPMethods([]string{http.MethodGet})(
				health.PingHandler)))
	mux.Handle(
		strings.Join([]string{"/v1/translations", "supported-locales"}, "/"),
		loggerMiddleware(
			allowHTTPMethods([]string{http.MethodGet})(
				authenticateApps([]string{"pocket-slate"})(
					translations.GetSupportedLocalesHandler))))
	mux.Handle(
		"/v1/translations",
		loggerMiddleware(
			allowHTTPMethods([]string{http.MethodPost})(
				authenticateApps([]string{"pocket-slate"})(
					translations.MakeTranslationHandler))))
	mux.Handle(
		"/",
		loggerMiddleware(
			allowHTTPMethods([]string{http.MethodGet})(
				NotFound)))

	log.Printf("Listening on %s...", serverAddress)

	err := http.ListenAndServe(serverAddress, mux)
	log.Fatal(err)
}
