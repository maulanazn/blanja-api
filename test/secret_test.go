package test

import (
	"fmt"
	"os"
	"testing"
	"userboilerplate-api/config"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDotenv(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	env := os.Getenv("CLD_URL")

	fmt.Println(env)
}

func TestGetCLDURL(t *testing.T) {
	config := config.GetConfig()
	assert.Equal(t, "localhost", config.GetString("DB_HOST"))
}
