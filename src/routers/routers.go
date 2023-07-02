package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	health.HealthRoutes(mux, loggerMiddleware)
	translations.TranslationsRoutes(mux, loggerMiddleware)
	mux.Handle("/", loggerMiddleware(http.HandlerFunc(notFound)))

	log.Printf("Listening on %s...", serverAddress)

	err := http.ListenAndServe(serverAddress, mux)
	log.Fatal(err)
}
