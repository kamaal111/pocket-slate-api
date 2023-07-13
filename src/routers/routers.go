package routers

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/kamaal111/pocket-slate-api/src/docs"
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

	basePath := "api/v1"
	docs.SwaggerInfo.BasePath = basePath

	health.Router(engine, basePath)
	translations.Router(engine, basePath)

	engine.NoRoute(NotFound)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Run(serverAddress)
}
