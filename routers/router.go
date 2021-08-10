package routers

import (
	"net/http"
	"time"

	v1 "github.com/chiehting/apiGo-template/routers/api/v1"
	middleware "github.com/chiehting/apiGo-template/routers/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter is setting up routing
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLFiles("public/index.html")
	r.Static("/asset", "public/asset")
	r.GET("/favicon.ico", getFavicon)
	r.GET("/", getIndex)

	r.POST("/api/v1/user/register", v1.UserRegister)
	r.POST("/api/v1/user/signin", v1.UserSignIn)

	user := r.Group("/api/v1/user")
	user.Use(middleware.JWT())
	{
		user.GET("", v1.UserInfo)
		user.POST("/signout", v1.UserSignOut)
	}

	serviceGroup := r.Group("/service")
	serviceGroup.Use()
	{
		serviceGroup.GET("/healthcheck", healthcheck)
	}

	return r
}

func getFavicon(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	// c.File("public/asset/images/favicon.png")
	c.Status(http.StatusNoContent)
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func healthcheck(c *gin.Context) {
	message := "pong"
	c.JSON(http.StatusOK, gin.H{"time": time.Now().UTC().Format(time.RFC3339), "message": message})
}
