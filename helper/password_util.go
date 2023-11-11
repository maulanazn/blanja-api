package helper

import (
	"entity"
	"os"
	"request"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
	email string
}

var loginRequest request.LoginRequest
var claims *Claims = &Claims{
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: TOKEN_EXP,
	},
	email: loginRequest.Email,
}

const TOKEN_EXP = int64(30 * time.Minute)

func ComparePasswords(hashedPwd, plainPwd []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return err
	}

	return nil
}

func GenerateToken(users entity.Users, data interface{}) string {
	os.Setenv("JWT_KEY", "tes123kjsdf0223j")

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	PanicIfError(err)

	return token
}

func DecodeToken(token string) string {
	jwt.ParseWithClaims(token, claims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	return claims.email
}
