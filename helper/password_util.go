package helper

import (
	"entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPwd, plainPwd []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return err
	}

	return nil
}

func GenerateToken(users entity.Users, data interface{}) string {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": users.Id,
		"email":   users.Email,
		"exp":     time.Now().Add(time.Duration(time.Now().UTC().Day())),
	})
	token, err := generateToken.SignedString(data)
	PanicIfError(err)

	return token
}
