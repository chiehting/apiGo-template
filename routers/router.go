package routers

import (
	"net/http"

	"github.com/chiehting/go-template/pkg/log"
	"github.com/chiehting/go-template/routers/api"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLFiles("public/index.html")
	r.Static("/asset", "public/asset")

	r.GET("/favicon.ico", getFavicon)
	r.GET("/", getIndex)

	apiGroup := r.Group("/api")
	apiGroup.GET("/healthcheck", api.Healthcheck)

	return r
}

func getFavicon(c *gin.Context) {
	// favicon.icon return 204
	c.Status(http.StatusNoContent)
}

func getIndex(c *gin.Context) {
	log.Info("ininini")
	c.HTML(http.StatusOK, "index.html", nil)
}
