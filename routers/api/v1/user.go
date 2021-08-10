package v1

import (
	"net/http"

	"github.com/chiehting/apiGo-template/pkg/log"
	"github.com/chiehting/apiGo-template/pkg/statuscode"
	"github.com/chiehting/apiGo-template/services/users"
	"github.com/gin-gonic/gin"
)

// UserRegister is user register account
func UserRegister(c *gin.Context) {
	var user users.User
	if err := c.BindJSON(&user); err != nil {
		log.Error(err)
	}

	var statusCode = user.Register()
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": statuscode.GetMsg(statusCode),
	})
}

// UserSignIn is user sign in
func UserSignIn(c *gin.Context) {
	var user users.User
	if err := c.BindJSON(&user); err != nil {
		log.Error(err)
	}

	var statusCode = user.SignIn(c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": statuscode.GetMsg(statusCode),
		"result":  map[string]interface{}{"token": user.Token},
	})
}

// UserInfo is get user information
func UserInfo(c *gin.Context) {
	var UID, _ = c.Get("UID")

	var user users.User
	user.ID = UID.(int)

	var userCache = user.GetCache()
	userCache["Token"] = "secret"
	c.JSON(http.StatusOK, gin.H{
		"code":    statuscode.Success,
		"message": statuscode.GetMsg(statuscode.Success),
		"result":  userCache,
	})
}

// UserSignOut is user sign out
func UserSignOut(c *gin.Context) {
	UID, _ := c.Get("UID")

	var user users.User
	user.ID = UID.(int)

	var statusCode = user.SignOut()
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": statuscode.GetMsg(statusCode),
	})
}
