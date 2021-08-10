package middleware

import (
	"net/http"
	"strings"

	"github.com/chiehting/apiGo-template/pkg/jwt"
	"github.com/chiehting/apiGo-template/pkg/statuscode"
	"github.com/chiehting/apiGo-template/services/users"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]
		if len(token) < 116 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": statuscode.AuthMiss,
				"msg":  statuscode.GetMsg(statuscode.AuthMiss),
			})
			c.Abort()
			return
		}

		var payload, err = jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": statuscode.AuthFail,
				"msg":  statuscode.GetMsg(statuscode.AuthFail),
			})
			c.Abort()
			return
		}

		var user users.User
		user.ID = int(payload["sub"].(float64))

		var userCache = user.GetCache()
		if token != userCache["Token"] || c.ClientIP() != userCache["SignInIP"] {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": statuscode.AuthFail,
				"msg":  statuscode.GetMsg(statuscode.AuthFail),
			})
			c.Abort()
			return
		}

		c.Set("UID", user.ID)
		c.Next()
	}
}
