package config_test

import (
	"testing"
	"util"

	"github.com/stretchr/testify/assert"
)

func TestPostgreSQL(t *testing.T) {
  viper := util.LoadConfig("../", "blanja.yaml", "yaml")

  assert.Equal(t, "localhost", viper.GetString("database.dbhost"))
}
