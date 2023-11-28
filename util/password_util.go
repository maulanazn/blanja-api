package util

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

func ComparePasswords(hashedPwd, plainPwd []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return err
	}

	return nil
}

func GenerateToken(userid string, data interface{}) string {
	const TOKEN_EXP = int64(30 * time.Minute)

	type Claims struct {
		jwt.StandardClaims
		Id string
	}

	var viperconfig *viper.Viper = LoadConfig(".", "blanja.yaml", "yaml")
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TOKEN_EXP,
		},
		Id: userid,
	})

	token, err := generateToken.SignedString([]byte(viperconfig.GetString("secret.jwtkey")))
	PanicIfError(err)

	return token
}

func DecodeToken(token string) string {
	strings.Split(token, " ")
	var viperconfig *viper.Viper = LoadConfig(".", "blanja.yaml", "yaml")

	type NewClaims struct {
		jwt.StandardClaims
		Id string
	}

	claims := &NewClaims{}
	_, err := jwt.ParseWithClaims(token[7:], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viperconfig.GetString("secret.jwtkey")), nil
	})

	if err != nil {
		log.Println("failed to parse")
	}

	return claims.Id
}
