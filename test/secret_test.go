package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDotenv(t *testing.T) {
	env := os.Getenv("CLD_URL")

	fmt.Println(env)
}

func TestGetCLDURL(t *testing.T) {
	assert.Equal(t, "localhost", os.Getenv("DB_HOST"))
}
