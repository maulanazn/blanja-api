package repository

import (
	"config"
	"entity"
	"log"
)

func InitDBPostgreSQL() {
	if err := config.GetConnection().AutoMigrate(&entity.Users{}, &entity.Address{}, &entity.Wishlist{}); err != nil {
		log.Println(err.Error())
	}
}
