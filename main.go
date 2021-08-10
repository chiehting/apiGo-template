package main

import (
	"github.com/chiehting/apiGo-template/pkg/gin"
	"github.com/chiehting/apiGo-template/pkg/log"
	"github.com/chiehting/apiGo-template/services"
)

func init() {
	services.Service.Init()
}

func main() {
	server := gin.HTTPServer()
	log.Infof("http server listening port %s", server.Addr)
	log.Panic(server.ListenAndServe())
}
