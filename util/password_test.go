package util_test

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"testing"
	"util"
)

func extractToken(token string) string {
	var viperconfig *viper.Viper = util.LoadConfig("./../", "blanja.yaml", "yaml")

	type NewClaims struct {
		jwt.StandardClaims
		Id string
	}

	claims := &NewClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viperconfig.GetString("secret.jwtkey")), nil
	})

	if err != nil {
		log.Println("failed to parse")
	}

	return claims.Id
}

func TestExtractClaimsJWTToken(t *testing.T) {
	result := extractToken("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MDAwMDAwMDAwMDAsIklkIjoiMzM5OWQ3NDktZTI2ZS00YjQzLTk3NGUtYTFmMWNlY2RjNGQ4In0.bDOvojvwsNr3eyGXJP5zlK5B7aHK3yLVXhFOSeSyeL2vtlx4q0FvoNr3Kg82EYXG4mLWiuNKi_dVr_nr5U8cTw")
	t.Log(result)
}
