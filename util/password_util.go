package util

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
  Id string
}


const TOKEN_EXP = int64(30 * time.Minute)

func ComparePasswords(hashedPwd, plainPwd []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return err
	}

	return nil
}

// var viperconfig *viper.Viper = LoadConfig("./../", "blanja.yaml", "yaml")

func GenerateToken(userid string, data interface{}) string {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
	  StandardClaims: jwt.StandardClaims{
		  ExpiresAt: TOKEN_EXP,
	  },
    Id: userid,
  })
	// token, err := generateToken.SignedString([]byte(viperconfig.GetString("secret.jwtkey")))
	// PanicIfError(err)
	token, err := generateToken.SignedString([]byte("testing123"))
	PanicIfError(err)

	return token
}

func DecodeToken(token string, req *http.Request) string {
  userid, err := req.Cookie("USR_ID")
  PanicIfError(err)
  var claims *Claims = &Claims{
	  StandardClaims: jwt.StandardClaims{
		  ExpiresAt: TOKEN_EXP,
	  },
    Id: userid.Value,
  }

	jwt.ParseWithClaims(token, claims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte("testing123"), nil
	})

	return claims.Id
}
