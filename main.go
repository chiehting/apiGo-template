package main

import (
	"github.com/chiehting/go-template/pkg/database"
	"github.com/chiehting/go-template/pkg/gin"
	"github.com/chiehting/go-template/pkg/log"
)

func init() {
	err := database.RunMigration()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	server := gin.HttpServer()
	log.Panic(server.ListenAndServe())
}
