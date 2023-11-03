package helper

import (
	"entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(users entity.Users, data interface{}) string {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": users.Username,
		"email":    users.Email,
		"exp":      time.Now().Add(time.Duration(time.Now().UTC().Day())),
	})
	token, err := generateToken.SignedString(data)
	PanicIfError(err)

	return token
}
