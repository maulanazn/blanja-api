package test

import (
	"fmt"
	"os"
	"testing"

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

func TestJwtDecode(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hdWxhbmF6bjE5QG1haWwuY29tIiwiZXhwIjoiMjAyMy0xMS0xMFQyMDo0ODo1NC41OTY3OTk1NTMrMDc6MDAiLCJ1c2VybmFtZSI6IiJ9.P4ytANL4AtUhLV22dSaPJ39zZItB5WRqQVkTp0pAf0g"
	mapClaim := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, mapClaim, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
	if err != nil {
		panic(err)
	}

	for key, val := range mapClaim {
		fmt.Println(key, val)
	}
}
