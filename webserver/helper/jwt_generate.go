package helper

import (
	"belanjabackend/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(customer entity.Customer, data interface{}) string {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": customer.Username,
		"email":    customer.Email,
		"exp":      time.Now().Add(time.Duration(time.Now().UTC().Day())),
	})
	token, err := generateToken.SignedString(data)
	PanicIfError(err)

	return token
}
