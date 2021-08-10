package users

import (
	"errors"
	"net/mail"
	"time"

	"github.com/chiehting/apiGo-template/models"
	"github.com/chiehting/apiGo-template/pkg/convert"
	"github.com/chiehting/apiGo-template/pkg/jwt"
	"github.com/chiehting/apiGo-template/pkg/log"
	"github.com/chiehting/apiGo-template/pkg/statuscode"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// User is user information
type User struct {
	models.User
}

const (
	active = iota + 1
	blocked
)

// Register is create user account
func (user *User) Register() int {
	if ok := user.verifyFormat(); !ok {
		return statuscode.UserRegisterFormatError
	}

	user.State = active
	user.EncryptedPassword, _ = bcrypt.GenerateFromPassword(convert.String2Bytes(user.Password), 8)
	if err := user.Create(); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return statuscode.UserRegisterDuplicate
		}
		log.Error(err)
		return statuscode.Error
	}

	return statuscode.Success
}

// SignIn is to check the account is exist and set the account information
func (user *User) SignIn(signInIP string) int {
	if isEmail(user.Account) {
		user.Email = user.Account
	} else {
		user.Username = user.Account
	}

	user.State = active
	if err := user.Get(); err != nil {
		return statuscode.UserNotFound
	}

	if ok := user.checkPassword(); !ok {
		return statuscode.UserSignInFail
	}

	user.SignInAt = time.Now()
	user.SignInIP = signInIP
	user.SignInCount = user.SignInCount + 1
	if err := user.SetSignIn(); err != nil {
		log.Error(err)
		return statuscode.Error
	}

	user.Token, _ = jwt.GenerateToken(user.ID)
	user.SetCache()
	user.SetExpireTime(time.Hour * 24)
	return statuscode.Success
}

// SignOut is to remove account cache
func (user *User) SignOut() int {
	if err := user.DelCache(); err != nil {
		return statuscode.Error
	}
	return statuscode.Success
}

func (user *User) checkPassword() bool {
	err := bcrypt.CompareHashAndPassword(user.EncryptedPassword, convert.String2Bytes(user.Password))
	return err == nil
}

func (user *User) verifyFormat() bool {
	if !isEmail(user.Email) {
		return false
	}
	return true
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
