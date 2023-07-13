package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
//
//	@Summary	Checks if server is up or down
//	@Schemes
//	@Description	Pings the server
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	pingResponse
//	@Router			/health [get]
func pingHandler(context *gin.Context) {
	context.JSON(http.StatusOK, pingResponse{Message: "pong"})
}
