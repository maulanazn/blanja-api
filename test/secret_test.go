package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestDotenv(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	env := os.Getenv("CLD_URL")

	fmt.Println(env)
}
