package repository

import (
	"config"
	"entity"
)

func InitDBPostgreSQL() {
	config.GetConnection().AutoMigrate(&entity.Users{}, &entity.Address{})
}
