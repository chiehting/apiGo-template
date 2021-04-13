package gin

import (
	"net/http"

	"github.com/chiehting/go-template/pkg/config"
	"github.com/chiehting/go-template/pkg/log"
	"github.com/chiehting/go-template/routers"
	"github.com/gin-gonic/gin"
)

func HttpServer() *http.Server {
	cfg := config.GetServerConfig()
	gin.SetMode(cfg.RunMode)
	routersInit := routers.InitRouter()
	server := &http.Server{
		Addr:           cfg.HTTPPort,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		Handler:        routersInit,
	}

	log.Infof("http server listening port %s", cfg.HTTPPort)
	return server
}
