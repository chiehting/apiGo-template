package jwt

import (
	"fmt"
	"time"

	"github.com/chiehting/apiGo-template/pkg/config"
	"github.com/golang-jwt/jwt"
)

var _cfg = config.GetApplication()
var _hmacSampleSecret = _cfg.Name + _cfg.Version

// GenerateToken is generate the user token
func GenerateToken(sub int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
	})

	if tokenString, err := token.SignedString([]byte(_hmacSampleSecret)); err == nil {
		return tokenString, nil
	}

	return "", jwt.ErrSignatureInvalid
}

// ParseToken parsing token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(_hmacSampleSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
