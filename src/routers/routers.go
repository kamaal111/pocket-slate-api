package routers

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

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

	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(jsonMiddleware())

	health.Router(engine)
	translations.Router(engine)

	engine.Run(serverAddress)

	// mux.Handle(
	// 	strings.Join([]string{"/v1/translations", "supported-locales"}, "/"),
	// 	loggerMiddleware(
	// 		allowHTTPMethods([]string{http.MethodGet})(
	// 			authenticateApps([]string{"pocket-slate"})(
	// 				translations.GetSupportedLocalesHandler))))
	// mux.Handle(
	// 	"/",
	// 	loggerMiddleware(
	// 		allowHTTPMethods([]string{http.MethodGet})(
	// 			NotFound)))

	// log.Printf("Listening on %s...", serverAddress)

	// err := http.ListenAndServe(serverAddress, mux)
	// log.Fatal(err)
}
