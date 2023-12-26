package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) registerPingRoutes() {
	handler.api.GET("/ping", handleGetPing)
}

func handleGetPing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Live and running!!")
}
