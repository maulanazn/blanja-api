package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestDotenv(t *testing.T) {
	env := os.Getenv("CLD_URL")

	fmt.Println(env)
}

func TestGetCLDURL(t *testing.T) {
	assert.Equal(t, "localhost", os.Getenv("DB_HOST"))
}

func TestGetJWT(t *testing.T) {
	assert.Equal(t, "", nil, "JWT_KEY is nil")
}

type Claims struct {
	jwt.StandardClaims
	userId string
	email  string
}

var user = struct {
	userId string
	email  string
}{
	userId: "apalah",
	email:  "apalah@apalah.sdd",
}

var claims *Claims = &Claims{
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: TOKEN_EXP,
	},
	userId: user.userId,
	email:  user.email,
}

const TOKEN_EXP = int64(30 * time.Minute)

func TestJwtEncode(t *testing.T) {
	os.Setenv("JWT_KEY", "tes123")

	datatoken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := datatoken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		panic(err)
	}

	jwt.ParseWithClaims(token, claims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	log.Println(token)
}

func TestJwtEnccode(t *testing.T) {
	os.Setenv("JWT_KEY", "tes123")

	jwt.ParseWithClaims("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MDAwMDAwMDAwMDB9.8eYiqL0GUM47IfdP_4qkj7r11_CCehbTQrEdGduuD1smgls5Jd4TsQ2M-3HefjHzyM7427iQWJVo2Bl0ef7beA", claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	fmt.Println(claims.userId)
}
