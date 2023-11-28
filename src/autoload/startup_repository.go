package autoload

import (
	"address"
	"config"
	"log"
	"users"
	"wishlist"
)

func InitDBPostgreSQL() {
	if err := config.GetConnection().AutoMigrate(&users.Users{}, &address.Address{}, &wishlist.Wishlist{}); err != nil {
		log.Println(err.Error())
	}
}
