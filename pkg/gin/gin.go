package gin

import (
	"net/http"

	"github.com/chiehting/apiGo-template/pkg/config"
	"github.com/chiehting/apiGo-template/routers"
	"github.com/gin-gonic/gin"
)

// HTTPServer setting up for the started http server
func HTTPServer() *http.Server {
	cfg := config.GetServer()
	gin.SetMode(cfg.RunMode)
	routersInit := routers.InitRouter()
	server := &http.Server{
		Addr:           cfg.HTTPPort,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		Handler:        routersInit,
	}

	return server
}
