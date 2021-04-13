package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Healthcheck 健康確認
func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"time": time.Now().UTC().Format(time.RFC3339), "message": "pong"})
}
