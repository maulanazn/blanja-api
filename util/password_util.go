package util

import (
	"entity"
	"request"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
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

var viperconfig *viper.Viper = LoadConfig("./../", "blanja.yaml", "yaml")

func GenerateToken(users entity.Users, data interface{}) string {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := generateToken.SignedString([]byte(viperconfig.GetString("secret.jwtkey")))
	PanicIfError(err)

	return token
}

func DecodeToken(token string) string {
	jwt.ParseWithClaims(token, claims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(viperconfig.GetString("secret.jwtkey")), nil
	})

	return claims.email
}
